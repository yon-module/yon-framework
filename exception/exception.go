package exception

import "github.com/yon-module/yon-framework/server/response"

func ExceptionHandler(err any) response.BaseResponse {
	if ok, res := notFoundException(err); ok {
		return res.Response()
	}

	if ok, res := badRequestException(err); ok {
		return res.Response()
	}

	if ok, res := intenalServerException(err); ok {
		return res.Response()
	}

	return response.ErrorResponse(response.ServerError, "Internal server error", err)
}
