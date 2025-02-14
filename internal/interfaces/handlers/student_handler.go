package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/internal/usecases/student"
	"strconv"
	"strings"
)

type StudentHandler struct {
	createStudentUseCase *student.CreateStudentUseCase
	getUserByIDUseCase   *student.GetStudentByIDUseCase
}

func NewStudentHandler(createStudentUseCase *student.CreateStudentUseCase, getUserByIDUseCase *student.GetStudentByIDUseCase) *StudentHandler {
	return &StudentHandler{
		createStudentUseCase: createStudentUseCase,
		getUserByIDUseCase:   getUserByIDUseCase,
	}
}

func (h *StudentHandler) Create(c *gin.Context) {
	var input student.CreateStudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	student, err := h.createStudentUseCase.Execute(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, student)
}

func (h *StudentHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id format",
		})
		return
	}

	student, err := h.getUserByIDUseCase.Execute(uint((id)))

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error fetching student",
		})
		return
	}

	c.JSON(http.StatusOK, student)

}
