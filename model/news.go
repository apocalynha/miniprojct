package model

import (
	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	UserID  uint   `json:"user_id" form:"user_id"`
	User    User   `json:"user" gorm:"foreignkey:UserID"` // Explicitly define the foreign key
	Tittle  string `json:"tittle" form:"tittle"`
	Content string `json:"content" form:"content"`
}
