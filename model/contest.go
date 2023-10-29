package model

import (
	"gorm.io/gorm"
	"time"
)

type Contest struct {
	gorm.Model
	ContestName string         `json:"contest_name"`
	ReqGender   string         `json:"req_gender"`
	ReqCategory string         `json:"req_category"`
	Details     string         `json:"details"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
