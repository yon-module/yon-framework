package server

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yon-module/yon-framework/database"
	"github.com/yon-module/yon-framework/logger"
	"github.com/yon-module/yon-framework/middleware"
)

type Server struct {
	Router *gin.Engine
}

func NewServer() *Server {
	logger.InitLogger()

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.LoggerMiddleware())

	server := &Server{Router: r}
	return server
}

func (s *Server) Start() {
	s.initConfig()

	port := "8080"
	if os.Getenv("yon.server.port") != "" {
		port = os.Getenv("yon.server.port")
	}

	logger.Log.Info().Msg("Gin server running use port " + port)
	s.Router.Run(":" + port)
}

func (s *Server) initConfig() {
	logger.Log.Info().Msg("Initial database")
	database.InitDB()
}
