package entities

import (
	"time"

	"gorm.io/gorm"
)

type KnowledgeSource struct {
	ID string `gorm:"type:uuid;primaryKey"`

	Name         string       `gorm:"size:255;not null"`
	Provider     string       `gorm:"size:50;not null"`
	ProviderType ProviderType `gorm:"size:20;not null"`

	Config      EncryptedJSON `gorm:"type:jsonb;not null"`
	UseRAG      bool          `gorm:"default:true"`
	RAGStrategy RAGStrategy   `gorm:"size:20"`
	RAGConfig   *JSONMap      `gorm:"type:jsonb"`

	Status     SourceStatus `gorm:"size:20;default:'active'"`
	LastSyncAt *time.Time
	LastError  *string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
