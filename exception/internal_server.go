package exception

import "github.com/yon-module/yon-framework/server/response"

type IntenalServerExceptionStruct struct {
	ErrorCode int
	Error     string
}

func NewIntenalServerExceptionStruct(code int, error string) IntenalServerExceptionStruct {
	return IntenalServerExceptionStruct{
		ErrorCode: code,
		Error:     error,
	}
}

func intenalServerException(err any) (bool, IntenalServerExceptionStruct) {
	if e, ok := err.(IntenalServerExceptionStruct); ok {
		return true, e
	}
	return false, IntenalServerExceptionStruct{}
}

func (n IntenalServerExceptionStruct) Response() response.BaseResponse {
	return response.ErrorResponse(
		n.ErrorCode,
		n.Error,
		nil,
	)
}
