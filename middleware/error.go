package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/yon-module/yon-framework/exception"
	"github.com/yon-module/yon-framework/logger"
	"github.com/yon-module/yon-framework/server/response"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			logger.Log.Error().Msgf("Error: %v", c.Errors.String())

			c.JSON(http.StatusInternalServerError, response.ErrorResponse(
				response.ServerError, "Internal Server Error", c.Errors.String(),
			))
			c.Abort()
		}
	}
}

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stack := string(debug.Stack())

				// Log readable format
				logger.Log.Error().
					Str("error", fmt.Sprintf("%v", err)).
					Msg(fmt.Sprintf("[PANIC RECOVERED]\nError: %v\nStacktrace:\n%s", err, stack))

				exception.ExceptionHandler(err).Json(c)
				c.Abort()
			}
		}()
		c.Next()
	}
}
