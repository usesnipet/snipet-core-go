package knowledge_source

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/usesnipet/snipet-core-go/internal/entity"
	http_server "github.com/usesnipet/snipet-core-go/internal/infra/http-server"
	"gorm.io/gorm"
)

type Controller struct {
	*http_server.BaseController
	service Service
}

func NewController(service Service, validate *validator.Validate) *Controller {
	return &Controller{
		BaseController: http_server.NewBaseController(validate),
		service:        service,
	}
}

func (c *Controller) RegisterRoutes(r chi.Router) {
	r.Route("/knowledge-sources", func(r chi.Router) {
		r.Get("/", c.list)
		r.Post("/", c.create)
		r.Get("/{id}", c.getByID)
		r.Patch("/{id}", c.update)
		r.Delete("/{id}", c.delete)
	})
}

func (c *Controller) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sources, err := c.service.FindAll(ctx)
	if err != nil {
		c.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	c.WriteJSON(w, http.StatusOK, sources)
}

func (c *Controller) getByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type params struct {
		ID string `param:"id" validate:"required,uuid"`
	}

	var p params
	if ok := c.DecodeParams(w, r, &p); !ok {
		return
	}

	source, err := c.service.FindByID(ctx, p.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.WriteJSONError(w, http.StatusNotFound, "knowledge source not found")
			return
		}
		c.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	c.WriteJSON(w, http.StatusOK, source)
}

func (c *Controller) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var dto CreateDTO
	if ok := c.DecodeBody(w, r, &dto); !ok {
		return
	}
	ent, err := c.service.Create(ctx, &dto)
	if err != nil {
		c.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	c.WriteJSON(w, http.StatusCreated, ent)
}

func (c *Controller) update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type params struct {
		ID string `param:"id" validate:"required,uuid"`
	}

	var p params
	if ok := c.DecodeParams(w, r, &p); !ok {
		return
	}

	var dto UpdateDTO
	if ok := c.DecodeBody(w, r, &dto); !ok {
		return
	}

	existing, err := c.service.FindByID(ctx, p.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.WriteJSONError(w, http.StatusNotFound, "knowledge source not found")
			return
		}
		c.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	applyUpdateDTOToEntity(&dto, existing)
	if err := c.service.Update(ctx, existing); err != nil {
		c.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	c.WriteJSON(w, http.StatusOK, existing)
}

func (c *Controller) delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type params struct {
		ID string `param:"id" validate:"required,uuid"`
	}

	var p params
	if ok := c.DecodeParams(w, r, &p); !ok {
		return
	}

	if err := c.service.Delete(ctx, p.ID); err != nil {
		c.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func applyUpdateDTOToEntity(dto *UpdateDTO, source *entity.KnowledgeSource) {
	if dto.Name != nil {
		source.Name = *dto.Name
	}
	if dto.Provider != nil {
		source.Provider = *dto.Provider
	}
	if dto.ProviderType != nil {
		source.ProviderType = entity.ProviderType(*dto.ProviderType)
	}
	if dto.Config != nil {
		source.Config = entity.EncryptedJSON(dto.Config)
	}
	if dto.UseRAG != nil {
		source.UseRAG = *dto.UseRAG
	}
	if dto.RAGStrategy != nil {
		source.RAGStrategy = entity.RAGStrategy(*dto.RAGStrategy)
	}
	if dto.RAGConfig != nil {
		m := entity.JSONMap(dto.RAGConfig)
		source.RAGConfig = &m
	}
	if dto.Status != nil {
		source.Status = entity.SourceStatus(*dto.Status)
	}
}
