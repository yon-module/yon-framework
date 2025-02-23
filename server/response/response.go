package response

type BaseResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
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
		Code:    ValidationErr,
		Message: "Validation failed",
		Errors:  errors,
	}
}
