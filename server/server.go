package server

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yon-module/yon-framework/database"
	"github.com/yon-module/yon-framework/logger"
	"github.com/yon-module/yon-framework/middleware"
	"github.com/yon-module/yon-framework/server/response"
)

type Server struct {
	Router      *gin.Engine
	RouterGroup *gin.RouterGroup
}

func NewServer() *Server {
	_ = godotenv.Load()
	logger.InitLogger()

	if os.Getenv("yon.server.env") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.RecoveryMiddleware())

	r.NoMethod(func(context *gin.Context) {
		response.ErrorResponse(
			response.MethodNotAllowed,
			"Method not found "+context.Request.Method,
			nil,
		).Json(context)
		return
	})

	r.NoRoute(func(context *gin.Context) {
		response.ErrorResponse(
			response.NotFound,
			"Route Not Found with method "+context.Request.Method,
			nil,
		).Json(context)
		return
	})

	contextPath := os.Getenv("yon.server.context.path")
	if contextPath == "" {
		contextPath = "/"
	}

	logger.Log.Info().Msg("Connect to Yon " + contextPath)

	def := r.Group(contextPath)
	def.GET("/ping", func(ctx *gin.Context) {
		response.SuccessResponse("Success ping", os.Getenv("yon.server.appName")).Json(ctx)
	})

	r.HandleMethodNotAllowed = true
	server := &Server{Router: r, RouterGroup: def}
	return server
}

func (s *Server) Start() {
	s.initConfig()

	port := "8080"
	if os.Getenv("yon.server.port") != "" {
		port = os.Getenv("yon.server.port")
	}

	logger.Log.Info().Msg("Yon server running use port " + port)
	_ = s.Router.Run(":" + port)
}

func (s *Server) initConfig() {
	logger.Log.Info().Msg("Initial database")
	database.InitDB()
}
