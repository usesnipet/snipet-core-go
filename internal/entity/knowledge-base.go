package entity

import (
	"time"

	"gorm.io/gorm"
)

type KnowledgeBase struct {
	ID          string         `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string         `gorm:"size:255;not null;uniqueIndex" json:"name"`
	Description *string        `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`

	Sources []KnowledgeSource `gorm:"many2many:knowledge_base_sources" json:"sources"`
}
