package entities

import "gorm.io/gorm"

type Knowledge struct {
	gorm.Model

	Name        string `json:"name" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:varchar(255)"`
}
