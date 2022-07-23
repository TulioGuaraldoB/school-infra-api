package student

import (
	"time"

	"github.com/TulioGuaraldoB/school-report/db/entity"
	"github.com/TulioGuaraldoB/school-report/util/pagination"
)

type ReportResponse struct {
	ID       uint            `json:"report_id"`
	Message  string          `json:"message"`
	Location string          `json:"location"`
	Student  StudentResponse `json:"student"`
}

type ReportRequest struct {
	Message   string `json:"message"`
	Location  string `json:"location"`
	StudentID uint   `json:"student_id"`
}

type StudentResponse struct {
	Name  string `json:"name"`
	Class string `json:"class"`
	Email string `json:"email"`
}

type StudentRequest struct {
	Name  string `json:"name"`
	Class string `json:"class"`
	Email string `json:"email"`
}

type RequestAll struct {
	Pagination pagination.Pagination
	Entity     entity.Student
}

type Credentials struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Username string `json:"user_name"`
	Password string `json:"password"`
}

func EntityToResponse(student *entity.Student, res *StudentResponse) {
	*res = StudentResponse{
		Name:  student.Name,
		Class: student.Class,
		Email: student.Email,
	}
}

func RequestToCreate(req *StudentRequest, student *entity.Student) {
	currentFormated := time.Now().Format("2017-07-09")
	now, _ := time.Parse("09-07-2017", currentFormated)

	*student = entity.Student{
		Name:      req.Name,
		Class:     req.Class,
		Email:     req.Email,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func ReportToResponse(report *entity.StudentReport, res *ReportResponse) {
	*res = ReportResponse{
		ID:       report.ID,
		Message:  report.Message,
		Location: report.Location,
		Student: StudentResponse{
			Name:  report.Student.Name,
			Class: report.Student.Class,
			Email: report.Student.Email,
		},
	}
}

func RequestToMessageReport(req *ReportRequest, report *entity.StudentReport) {
	currentFormated := time.Now().Format("2017-07-09")
	now, _ := time.Parse("09-07-2017", currentFormated)

	*report = entity.StudentReport{
		Message:   req.Message,
		Location:  req.Location,
		StudentID: req.StudentID,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func RequestToUser(req *UserRequest, user *entity.User) {
	*user = entity.User{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}
}
