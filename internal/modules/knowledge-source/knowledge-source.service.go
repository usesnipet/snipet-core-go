package knowledge_source

import (
	"context"

	"github.com/usesnipet/snipet-core-go/internal/entity"
)

type Service interface {
	Create(ctx context.Context, source *CreateDTO) (*entity.KnowledgeSource, error)
	FindByID(ctx context.Context, id string) (*entity.KnowledgeSource, error)
	FindAll(ctx context.Context) ([]entity.KnowledgeSource, error)
	Update(ctx context.Context, source *entity.KnowledgeSource) error
	Delete(ctx context.Context, id string) error
}

type service struct {
	repository Repository
}

func (s *service) Create(ctx context.Context, dto *CreateDTO) (*entity.KnowledgeSource, error) {
	source := s.createDTOToEntity(dto)
	err := s.repository.Create(ctx, source)
	return source, err
}

func (s *service) FindByID(ctx context.Context, id string) (*entity.KnowledgeSource, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *service) FindAll(ctx context.Context) ([]entity.KnowledgeSource, error) {
	return s.repository.FindAll(ctx)
}

func (s *service) Update(ctx context.Context, source *entity.KnowledgeSource) error {
	return s.repository.Update(ctx, source)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}

func (s *service) createDTOToEntity(dto *CreateDTO) *entity.KnowledgeSource {
	source := &entity.KnowledgeSource{
		Name:         dto.Name,
		Provider:     dto.Provider,
		ProviderType: entity.ProviderType(dto.ProviderType),
		Config:       entity.EncryptedJSON(dto.Config),
		RAGStrategy:  entity.RAGStrategy(dto.RAGStrategy),
		Status:       entity.SourceStatus(dto.Status),
	}
	if dto.UseRAG != nil {
		source.UseRAG = *dto.UseRAG
	} else {
		source.UseRAG = true
	}
	if dto.Status == "" {
		source.Status = entity.SourceActive
	}
	if dto.RAGConfig != nil {
		m := entity.JSONMap(dto.RAGConfig)
		source.RAGConfig = &m
	}
	return source
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}
