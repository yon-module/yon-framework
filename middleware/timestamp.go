package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yon-module/yon-framework/server/response"
)

func ValidateTimestamp(maxSkew time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		tsStr := c.GetHeader("X-Timestamp")
		if tsStr == "" {
			response.ErrorResponse(response.Unauthorized, "Missing X-Timestamp header", nil).Json(c)
			c.Abort()
			return
		}

		tsInt, err := strconv.ParseInt(tsStr, 10, 64)
		if err != nil {
			response.ErrorResponse(response.Unauthorized, "Invalid timestamp", nil).Json(c)
			c.Abort()
			return
		}

		reqTime := time.Unix(tsInt, 0)
		now := time.Now()

		diff := now.Sub(reqTime)
		if diff < 0 {
			diff = -diff
		}

		if diff > maxSkew {
			response.ErrorResponse(response.Unauthorized, "Expired request", nil).Json(c)
			c.Abort()
			return
		}

		c.Set("requestTime", reqTime)
		c.Next()
	}
}
