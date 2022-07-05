package student

import (
	"github.com/TulioGuaraldoB/school-report/db/entity"
)

type interfaceService interface {
	all(requestAll *RequestAll) ([]entity.Student, error)
	show(id uint) (*entity.Student, error)
	create(student *entity.Student) error
	allReports() ([]entity.StudentReport, error)
	createReport(report *entity.StudentReport) error
}

type service struct {
	repository interfaceRepository
}

func NewService(repository interfaceRepository) interfaceService {
	return &service{
		repository: repository,
	}
}

func (s *service) all(requestAll *RequestAll) ([]entity.Student, error) {
	return s.repository.all(requestAll)
}

func (s *service) show(id uint) (*entity.Student, error) {
	return s.repository.show(id)
}

func (s *service) create(student *entity.Student) error {
	return s.repository.create(student)
}

func (s *service) allReports() ([]entity.StudentReport, error) {
	return s.repository.allReports()
}

func (s *service) createReport(report *entity.StudentReport) error {
	return s.repository.createReport(report)
}
