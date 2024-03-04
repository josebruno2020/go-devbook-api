package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey,serializer:json"json:"id,omitempty"`
	Name      string         `json:"name,omitempty"binding:"required"`
	Nick      string         `json:"nick,omitempty"binding:"required"`
	Email     string         `json:"email,omitempty"binding:"required"`
	Password  string         `json:"-"binding:"required"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"json:"-"`
}
