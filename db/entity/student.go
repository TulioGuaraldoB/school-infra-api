package entity

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Class     string `gorm:"size:125"`
	Email     string `gorm:"size:255"`
	ReportID  int
	Report    StudentReport `gorm:"foreignKey:ReportID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
