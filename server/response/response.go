package response

import (
	"os"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Status    string `json:"status"`
	RequestID string `json:"requestId"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
	Errors    any    `json:"errors,omitempty"`
}

func SuccessResponse(message string, data any) BaseResponse {
	return BaseResponse{
		Status:  "success",
		Code:    Success,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(code int, message string, errors any) BaseResponse {
	if os.Getenv("yon.server.env") == "production" {
		errors = nil
	}
	return BaseResponse{
		Status:  "error",
		Code:    code,
		Message: message,
		Errors:  errors,
	}
}

func ValidationErrorResponse(errors any) BaseResponse {
	return BaseResponse{
		Status:  "fail",
		Code:    UnprocessableEntity,
		Message: "Validation failed",
		Errors:  errors,
	}
}

func (r BaseResponse) Json(ctx *gin.Context) {
	r.RequestID = ctx.GetString("requestid")
	ctx.JSON(getCustomCode(r.Code), r)
	ctx.Abort()
}

func (r BaseResponse) Build(ctx *gin.Context) {
	r.Json(ctx)
	ctx.Abort()
}
