package entities

type KnowledgeBaseSource struct {
	KnowledgeBaseID   string `gorm:"type:uuid;primaryKey"`
	KnowledgeSourceID string `gorm:"type:uuid;primaryKey"`
}
