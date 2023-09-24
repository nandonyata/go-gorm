package model

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Title      string         `json:"title"`
	UserID uint `json:"userId"`
}