package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `json:"name" gorm:"not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	Role      string `json:"role" gorm:"type:enum('user','admin');default:'user'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
