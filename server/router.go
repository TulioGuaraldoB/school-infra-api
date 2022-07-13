package server

import (
	"net/http"

	"github.com/TulioGuaraldoB/school-report/core/student"
	"github.com/TulioGuaraldoB/school-report/db"
	"github.com/gin-gonic/gin"
)

func GetRoutes() *gin.Engine {
	router := gin.Default()

	db := db.OpenConnection()

	studentRepository := student.NewRepository(db)
	studentService := student.NewService(studentRepository)
	studentController := student.NewController(studentService)

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			student := v1.Group("student")
			{
				student.GET("", studentController.GetAll)
				student.GET(":id", studentController.GetById)
				student.POST("", studentController.CreateStudent)
				student.GET("report", studentController.GetAllReports)
			}

			report := student.Group(":id/report")
			{
				report.POST("", studentController.CreateReport)
			}
		}

		v1.POST("login", studentController.Login)
	}

	router.GET("health", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "Healthy!")
	})

	return router
}
