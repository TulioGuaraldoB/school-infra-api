package student

import (
	"github.com/TulioGuaraldoB/school-report/db/entity"
	"github.com/TulioGuaraldoB/school-report/util"
	"gorm.io/gorm"
)

type interfaceRepository interface {
	all(student *entity.Student, pagination util.Pagination) ([]entity.Student, error)
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

func (r *repository) all(student *entity.Student, pagination util.Pagination) ([]entity.Student, error) {
	students := []entity.Student{}

	offset := (pagination.Page - 1) * pagination.Limit
	query := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	result := query.Joins("Report").Where(student).Find(&students)

	if result.Error != nil {
		return nil, result.Error
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
