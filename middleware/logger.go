package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yon-module/yon-framework/logger"
)

const (
	HeaderRequestID   = "X-Request-ID"
	KeyRequestID      = "requestid"
	KeyRequestContext = "request-ctx"
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
					Str("requestId", ctx.GetString(KeyRequestID)).
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
			Str("requestId", ctx.GetString(KeyRequestID)).
			Str("appName", appName).
			Str("method", ctx.Request.Method).
			Str("path", ctx.Request.URL.Path).
			Int("status", ctx.Writer.Status()).
			Int("size", ctx.Writer.Size()).
			Dur("latency", time.Since(start)).
			Msg("Request processed")
	}
}

func LoggerRequestMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := ctx.Request.Header.Get(HeaderRequestID)
		if requestID == "" {
			requestID, _ = generateRequestID(8)
		}
		ctx.Set(KeyRequestID, requestID)
		ctx.Next()
	}
}

func generateRequestID(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
