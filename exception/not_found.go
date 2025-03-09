package exception

import "github.com/yon-module/yon-framework/server/response"

type NotFoundExceptionStruct struct {
	ErrorCode int
	Error     string
}

func NewNotFoundException(code int, error string) NotFoundExceptionStruct {
	return NotFoundExceptionStruct{
		ErrorCode: code,
		Error:     error,
	}
}

func notFoundException(err any) (bool, NotFoundExceptionStruct) {
	if e, ok := err.(NotFoundExceptionStruct); ok {
		return true, e
	}
	return false, NotFoundExceptionStruct{}
}

func (n NotFoundExceptionStruct) Response() response.BaseResponse {
	return response.ErrorResponse(
		n.ErrorCode,
		n.Error,
		nil,
	)
}
