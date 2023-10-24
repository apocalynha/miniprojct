package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Email     string
	Password  string
	Role      string `json:"role" gorm:"type:enum('customer','admin');default:'customer'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
