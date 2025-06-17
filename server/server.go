package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yon-module/yon-framework/logger"
	"github.com/yon-module/yon-framework/middleware"
	"github.com/yon-module/yon-framework/server/response"
	"github.com/yon-module/yon-framework/yonevent"
	"os"
)

type Server struct {
	Router      *gin.Engine
	RouterGroup *gin.RouterGroup
}

var routes []func(gr *gin.RouterGroup)

func NewServer() *Server {
	logger.Log.Info().Msg("Starting yon server...")
	if os.Getenv("yon.server.env") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	initialLoggerSentry(r)
	r.Use(middleware.LoggerRequestMiddleware())
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

	for _, v := range routes {
		v(def)
	}

	r.HandleMethodNotAllowed = true
	server := &Server{Router: r, RouterGroup: def}
	return server
}

func (s *Server) Start() {

	port := "8080"
	if os.Getenv("yon.server.port") != "" {
		port = os.Getenv("yon.server.port")
	}

	logger.Log.Info().Msg("Yon server running use port " + port)
	if err := s.Router.Run(":" + port); err != nil {
		yonevent.Emit("serverready")
	}
}

func AddRoute(handler func(gr *gin.RouterGroup)) {
	routes = append(routes, handler)
}

func AddRoutes(handlers ...func(gr *gin.RouterGroup)) {
	routes = append(routes, handlers...)
}

func GetRequestId(ctx *gin.Context) string {
	return ctx.GetString("requestid")
}
