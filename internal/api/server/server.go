package server

import (
	"github.com/gin-gonic/gin"
	"rest-api/internal/api/middleware"
	"rest-api/internal/api/routes"
	"rest-api/internal/interfaces/handlers"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	return &Server{
		router: gin.Default(),
	}
}

func (s *Server) SetupRoutes(studentHandler *handlers.StudentHandler) {

	s.router.Use(middleware.ErrorHandler())

	routes.SetupStudentRoutes(s.router, studentHandler)
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
