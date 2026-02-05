package knowledge_source

import (
	"context"

	"github.com/usesnipet/snipet-core-go/internal/entity"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, source *entity.KnowledgeSource) error
	FindByID(ctx context.Context, id string) (*entity.KnowledgeSource, error)
	FindAll(ctx context.Context) ([]entity.KnowledgeSource, error)
	Update(ctx context.Context, source *entity.KnowledgeSource) error
	Delete(ctx context.Context, id string) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(ctx context.Context, source *entity.KnowledgeSource) error {
	return r.db.WithContext(ctx).Create(source).Error
}

func (r *repository) FindByID(ctx context.Context, id string) (*entity.KnowledgeSource, error) {
	var source entity.KnowledgeSource
	err := r.db.WithContext(ctx).First(&source, "id = ?", id).Error
	return &source, err
}
func (r *repository) FindAll(ctx context.Context) ([]entity.KnowledgeSource, error) {
	var sources []entity.KnowledgeSource
	err := r.db.WithContext(ctx).Find(&sources).Error
	return sources, err
}

func (r *repository) Update(ctx context.Context, source *entity.KnowledgeSource) error {
	err := r.db.WithContext(ctx).Save(source).Error
	return err
}

func (r *repository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&entity.KnowledgeSource{}, "id = ?", id).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}
