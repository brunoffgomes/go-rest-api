package routes

import (
	"github.com/gin-gonic/gin"
	"rest-api/internal/interfaces/handlers"
)

func SetupStudentRoutes(router *gin.Engine, studentHandler *handlers.StudentHandler) {
	api := router.Group("/api/v1")
	{
		students := api.Group("/students")
		{
			students.POST("/", studentHandler.Create)
			students.GET("/:id", studentHandler.GetByID)
		}
	}

}
