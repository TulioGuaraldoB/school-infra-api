package student

import (
	"github.com/TulioGuaraldoB/school-report/db/entity"
	"github.com/TulioGuaraldoB/school-report/util/jwt"
)

type interfaceService interface {
	all(requestAll *RequestAll) ([]entity.Student, error)
	getStudent(id uint) (*entity.Student, error)
	create(student *entity.Student) error
	allReports() ([]entity.StudentReport, error)
	createReport(report *entity.StudentReport) error
	getByCredentials(credentials Credentials) (*string, error)
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

func (s *service) getStudent(id uint) (*entity.Student, error) {
	return s.repository.getStudent(id)
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

func (s *service) getByCredentials(credentials Credentials) (*string, error) {
	user, err := s.repository.getByCredentials(credentials)
	if err != nil {
		return nil, err
	}

	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return nil, err
	}

	return token, nil
}
