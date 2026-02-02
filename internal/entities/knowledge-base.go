package entities

import (
	"time"

	"gorm.io/gorm"
)

type KnowledgeBase struct {
	ID          string `gorm:"type:uuid;primaryKey"`
	Name        string `gorm:"size:255;not null;uniqueIndex"`
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Sources []KnowledgeSource `gorm:"many2many:knowledge_base_sources"`
}
