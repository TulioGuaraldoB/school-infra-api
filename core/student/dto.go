package student

import (
	"time"

	"github.com/TulioGuaraldoB/school-report/db/entity"
)

type ReportResponse struct {
	ID       uint   `json:"report_id"`
	Message  string `json:"message"`
	Location string `json:"location"`
}

type ReportRequest struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type StudentResponse struct {
	Name   string         `json:"name"`
	Class  string         `json:"class"`
	Email  string         `json:"email"`
	Report ReportResponse `json:"report"`
}

type StudentRequest struct {
	Name  string `json:"name"`
	Class string `json:"class"`
	Email string `json:"email"`
}

func EntityToResponse(student *entity.Student, res *StudentResponse) {
	*res = StudentResponse{
		Name:  student.Name,
		Class: student.Class,
		Email: student.Email,
		Report: ReportResponse{
			ID:       student.Report.ID,
			Message:  student.Report.Message,
			Location: student.Report.Location,
		},
	}
}

func RequestToCreate(req *StudentRequest, student *entity.Student) {
	currentFormated := time.Now().Format("2017-07-09")
	now, _ := time.Parse("09-07-2017", currentFormated)

	*student = entity.Student{
		Name:      req.Name,
		Class:     req.Class,
		Email:     req.Email,
		ReportID:  1,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func RequestToMessageReport(req *ReportRequest, report *entity.StudentReport) {
	currentFormated := time.Now().Format("2017-07-09")
	now, _ := time.Parse("09-07-2017", currentFormated)

	*report = entity.StudentReport{
		Message:   req.Message,
		Location:  req.Location,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
