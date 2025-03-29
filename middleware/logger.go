package middleware

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yon-module/yon-framework/logger"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		appName := os.Getenv("yon.server.appName")
		if appName == "" {
			appName = "yon-framework"
		}

		errors := ctx.Errors
		if len(errors) > 0 {
			for _, err := range errors {
				logger.Log.Error().
					Str("appName", appName).
					Str("method", ctx.Request.Method).
					Str("path", ctx.Request.URL.Path).
					Int("status", ctx.Writer.Status()).
					Int("size", ctx.Writer.Size()).
					Dur("latency", time.Since(start)).
					Str("error", err.Error()).
					Msg("Request Failed")
			}
			return
		}

		logger.Log.Info().
			Str("appName", appName).
			Str("method", ctx.Request.Method).
			Str("path", ctx.Request.URL.Path).
			Int("status", ctx.Writer.Status()).
			Int("size", ctx.Writer.Size()).
			Dur("latency", time.Since(start)).
			Msg("Request processed")
	}
}
