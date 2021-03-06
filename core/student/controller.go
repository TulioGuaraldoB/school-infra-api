package student

import (
	"net/http"
	"strconv"

	"github.com/TulioGuaraldoB/school-report/db/entity"
	"github.com/TulioGuaraldoB/school-report/util/pagination"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type controller struct {
	service interfaceService
}

func NewController(service interfaceService) controller {
	return controller{
		service: service,
	}
}

func (c *controller) GetAll(ctx *gin.Context) {
	pagination := pagination.PaginationRequest(ctx)
	student := entity.Student{}

	reqAll := RequestAll{
		Pagination: pagination,
		Entity:     student,
	}

	students, err := c.service.all(&reqAll)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	studentsRes := []StudentResponse{}
	studentRes := StudentResponse{}

	for _, student := range students {
		EntityToResponse(&student, &studentRes)
		studentsRes = append(studentsRes, studentRes)

		reqAll.Pagination.Rows = studentsRes
	}

	ctx.IndentedJSON(http.StatusOK, reqAll.Pagination)
}

func (c *controller) GetById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	student, err := c.service.getStudent(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	studentRes := StudentResponse{}
	EntityToResponse(student, &studentRes)

	ctx.IndentedJSON(http.StatusOK, studentRes)
}

func (c *controller) CreateStudent(ctx *gin.Context) {
	studentReq := StudentRequest{}
	if err := ctx.ShouldBindJSON(&studentReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	student := entity.Student{}
	RequestToCreate(&studentReq, &student)

	if err := c.service.create(&student); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "student created successfully!",
		"student": studentReq,
	})
}

func (c *controller) GetAllReports(ctx *gin.Context) {
	reports, err := c.service.allReports()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	reportRes := ReportResponse{}
	reportsRes := []ReportResponse{}

	for _, report := range reports {
		ReportToResponse(&report, &reportRes)
		reportsRes = append(reportsRes, reportRes)
	}

	ctx.IndentedJSON(http.StatusOK, reportsRes)
}

func (c *controller) CreateReport(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	reportReq := ReportRequest{
		StudentID: uint(id),
	}
	if err := ctx.ShouldBindJSON(&reportReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	report := entity.StudentReport{}
	RequestToMessageReport(&reportReq, &report)

	if err := c.service.createReport(&report); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "report inserted successfully!",
		"student": reportReq,
	})
}

func (c *controller) Login(ctx *gin.Context) {
	signInCredentials := Credentials{}
	if err := ctx.ShouldBindJSON(&signInCredentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := c.service.getByCredentials(signInCredentials)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, token)
}

func (c *controller) Register(ctx *gin.Context) {
	userReq := UserRequest{}
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := entity.User{}
	RequestToUser(&userReq, &user)

	if err := c.service.createUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"user":    userReq,
		"message": "user registered successfully!",
	})
}
