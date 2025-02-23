package response

// ResponseCode mendefinisikan kode standar API
const (
	Success       = 1000
	BadRequest    = 1400
	Unauthorized  = 1401
	Forbidden     = 1403
	NotFound      = 1404
	Conflict      = 1409
	ValidationErr = 1422
	ServerError   = 1500
)
