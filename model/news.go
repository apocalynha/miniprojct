package model

import (
	"gorm.io/gorm"
	"time"
)

type News struct {
	gorm.Model
	UserID  uint   `json:"user_id" form:"user_id"`
	User    User   `json:"user" gorm:"foreignkey:UserID"` // Explicitly define the foreign key
	Tittle  string `json:"tittle" form:"tittle"`
	Content string `json:"content" form:"content"`
}

type NewsResponse struct {
	ID        uint      `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      User      `json:"user" ` // Explicitly define the foreign key
	Tittle    string    `json:"tittle" `
	Content   string    `json:"content"`
}

func (newsDB News) ResponseConvert() NewsResponse {
	var Response NewsResponse
	Response.ID = newsDB.ID
	Response.Tittle = newsDB.Tittle
	Response.Content = newsDB.Content

	return Response
}
