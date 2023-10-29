package model

import (
	"gorm.io/gorm"
)

type Contestant struct {
	gorm.Model
	UserID         uint    `json:"user_id" form:"user_id"`
	User           User    `json:"user" gorm:"foreignkey:UserID"` // Hubungkan dengan model User
	ContestID      uint    `json:"contest_id" form:"contest_id"`
	Contest        Contest `json:"contest" gorm:"foreignkey:ContestID"` // Hubungkan dengan model Contest
	ContestantName string  `json:"contestant_name" form:"contestant_name"`
	Gender         string  `json:"gender" gorm:"type:enum('Laki-laki','Perempuan')"`
	Age            int     `json:"age" form:"age"`
	Category       string  `json:"category" form:"category"`
}
