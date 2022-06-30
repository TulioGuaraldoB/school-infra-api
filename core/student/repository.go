package student

import (
	"github.com/TulioGuaraldoB/school-report/db/entity"
	"gorm.io/gorm"
)

type interfaceRepository interface {
	all() ([]entity.Student, error)
	show(id uint) (*entity.Student, error)
	create(student *entity.Student) error
	createReport(report *entity.StudentReport) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) interfaceRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) all() ([]entity.Student, error) {
	students := []entity.Student{}
	if err := r.db.Joins("Report").Find(&students).Error; err != nil {
		return nil, err
	}

	return students, nil
}

func (r *repository) show(id uint) (*entity.Student, error) {
	student := entity.Student{}
	if err := r.db.Joins("Report").First(&student, &id).Error; err != nil {
		return nil, err
	}

	return &student, nil
}

func (r *repository) create(student *entity.Student) error {
	if err := r.db.Create(&student).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) createReport(report *entity.StudentReport) error {
	if err := r.db.Create(&report).Error; err != nil {
		return err
	}

	return nil
}
