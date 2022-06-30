package entity

import (
	"time"

	"gorm.io/gorm"
)

type StudentReport struct {
	ID        uint   `gorm:"primaryKey"`
	Message   string `gorm:"325"`
	Location  string `gorm:"175"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
