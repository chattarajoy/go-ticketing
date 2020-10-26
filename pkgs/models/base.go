package models

import (
	"time"

	"gorm.io/gorm"
)

// base Model
type Model struct {
	ID        int            `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
