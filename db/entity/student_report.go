package entity

import (
	"time"

	"gorm.io/gorm"
)

type StudentReport struct {
	ID        uint   `gorm:"primaryKey"`
	Message   string `gorm:"325"`
	Location  string `gorm:"175"`
	StudentID uint
	Student   Student `gorm:"foreignKey:StudentID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
