package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yon-module/yon-framework/logger"
	"github.com/yon-module/yon-framework/server/response"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			logger.Log.Error().Str("Error:", c.Errors.String())

			c.JSON(http.StatusInternalServerError, response.ErrorResponse(
				response.ServerError, "Internal Server Error", c.Errors.String(),
			))
		}
	}
}

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Log.Error().Any("Panic occurred: %v", err)

				c.JSON(http.StatusInternalServerError, response.ErrorResponse(
					response.ServerError, "Internal Server Error", nil,
				))

				c.Abort()
			}
		}()
		c.Next()
	}
}
