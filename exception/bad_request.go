package exception

import "github.com/yon-module/yon-framework/server/response"

type BadRequestExceptionStruct struct {
	ErrorCode int
	Error     string
}

func NewBadRequestExceptionStruct(code int, error string) BadRequestExceptionStruct {
	return BadRequestExceptionStruct{
		ErrorCode: code,
		Error:     error,
	}
}

func badRequestException(err any) (bool, BadRequestExceptionStruct) {
	if e, ok := err.(BadRequestExceptionStruct); ok {
		return true, e
	}
	return false, BadRequestExceptionStruct{}
}

func (n BadRequestExceptionStruct) Response() response.BaseResponse {
	return response.ErrorResponse(
		n.ErrorCode,
		n.Error,
		nil,
	)
}
