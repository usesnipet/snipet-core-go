package entity

import (
	"time"

	"gorm.io/gorm"
)

type KnowledgeSource struct {
	ID string `gorm:"type:uuid;primaryKey" json:"id"`

	Name         string       `gorm:"size:255;not null" json:"name"`
	Provider     string       `gorm:"size:50;not null" json:"provider"`
	ProviderType ProviderType `gorm:"size:20;not null" json:"providerType"`

	Config      EncryptedJSON `gorm:"type:jsonb;not null" json:"config"`
	UseRAG      bool          `gorm:"default:true" json:"useRag"`
	RAGStrategy RAGStrategy   `gorm:"size:20" json:"ragStrategy"`
	RAGConfig   *JSONMap      `gorm:"type:jsonb" json:"ragConfig"`

	Status     SourceStatus `gorm:"size:20;default:'active'" json:"status"`
	LastSyncAt *time.Time   `json:"lastSyncAt"`
	LastError  *string      `json:"lastError"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
