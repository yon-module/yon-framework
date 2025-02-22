package server

import (
	"github.com/gin-gonic/gin"
	"github.com/username/goframework/database"
	"github.com/username/goframework/middleware"
	"log"
)

type Server struct {
	Router *gin.Engine
}

func NewServer() *Server {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	server := &Server{Router: r}
	return server
}

func (s *Server) Start(port string) {
	database.InitDB()
	log.Println("Server running on port", port)
	s.Router.Run(":" + port)
}
