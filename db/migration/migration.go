package migration

import (
	"github.com/TulioGuaraldoB/school-report/db/entity"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	db.AutoMigrate(&entity.StudentReport{})
	db.AutoMigrate(&entity.Student{})
	db.AutoMigrate(&entity.User{})
}
