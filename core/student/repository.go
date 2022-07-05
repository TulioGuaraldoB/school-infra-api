package student

import (
	"github.com/TulioGuaraldoB/school-report/db/entity"
	"gorm.io/gorm"
)

type interfaceRepository interface {
	all(requestAll *RequestAll) ([]entity.Student, error)
	show(id uint) (*entity.Student, error)
	create(student *entity.Student) error
	allReports() ([]entity.StudentReport, error)
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

func (r *repository) all(requestAll *RequestAll) ([]entity.Student, error) {
	students := []entity.Student{}

	offset := (requestAll.Pagination.Page - 1) * requestAll.Pagination.Limit
	query := r.db.Limit(requestAll.Pagination.Limit).Offset(offset).Order(requestAll.Pagination.Sort)

	result := query.
		Where(requestAll.Entity).
		Find(&students)

	if result.Error != nil {
		return nil, result.Error
	}

	return students, nil
}

func (r *repository) show(id uint) (*entity.Student, error) {
	student := entity.Student{}
	if err := r.db.First(&student, &id).Error; err != nil {
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

func (r *repository) allReports() ([]entity.StudentReport, error) {
	reports := []entity.StudentReport{}
	if err := r.db.Joins("Student").Find(&reports).Error; err != nil {
		return nil, err
	}

	return reports, nil
}

func (r *repository) createReport(report *entity.StudentReport) error {
	if err := r.db.Create(&report).Error; err != nil {
		return err
	}

	return nil
}
