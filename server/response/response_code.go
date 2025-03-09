package response

import "net/http"

// ResponseCode mendefinisikan kode standar API
const (
	Success           = 1000
	Created           = 1001
	Accepted          = 1002
	NoContent         = 1004
	MultipleChoices   = 1300
	MovedPermanently  = 1301
	Found             = 1302
	SeeOther          = 1303
	NotModified       = 1304
	TemporaryRedirect = 1307
	PermanentRedirect = 1308

	BadRequest                   = 1400
	Unauthorized                 = 1401
	PaymentRequired              = 1402
	Forbidden                    = 1403
	NotFound                     = 1404
	MethodNotAllowed             = 1405
	NotAcceptable                = 1406
	ProxyAuthRequired            = 1407
	RequestTimeout               = 1408
	Conflict                     = 1409
	Gone                         = 1410
	LengthRequired               = 1411
	PreconditionFailed           = 1412
	RequestEntityTooLarge        = 1413
	RequestURITooLong            = 1414
	UnsupportedMediaType         = 1415
	RequestedRangeNotSatisfiable = 1416
	ExpectationFailed            = 1417
	UnprocessableEntity          = 1422
	TooManyRequests              = 1429

	ServerError        = 1500
	NotImplemented     = 1501
	BadGateway         = 1502
	ServiceUnavailable = 1503
	GatewayTimeout     = 1504
)

var httpToCustomCode = map[int]int{}

func init() {
	statusMap := map[int]int{
		http.StatusOK:                           Success,
		http.StatusCreated:                      Created,
		http.StatusAccepted:                     Accepted,
		http.StatusNoContent:                    NoContent,
		http.StatusMultipleChoices:              MultipleChoices,
		http.StatusMovedPermanently:             MovedPermanently,
		http.StatusFound:                        Found,
		http.StatusSeeOther:                     SeeOther,
		http.StatusNotModified:                  NotModified,
		http.StatusTemporaryRedirect:            TemporaryRedirect,
		http.StatusPermanentRedirect:            PermanentRedirect,
		http.StatusBadRequest:                   BadRequest,
		http.StatusUnauthorized:                 Unauthorized,
		http.StatusPaymentRequired:              PaymentRequired,
		http.StatusForbidden:                    Forbidden,
		http.StatusNotFound:                     NotFound,
		http.StatusMethodNotAllowed:             MethodNotAllowed,
		http.StatusNotAcceptable:                NotAcceptable,
		http.StatusProxyAuthRequired:            ProxyAuthRequired,
		http.StatusRequestTimeout:               RequestTimeout,
		http.StatusConflict:                     Conflict,
		http.StatusGone:                         Gone,
		http.StatusLengthRequired:               LengthRequired,
		http.StatusPreconditionFailed:           PreconditionFailed,
		http.StatusRequestEntityTooLarge:        RequestEntityTooLarge,
		http.StatusRequestURITooLong:            RequestURITooLong,
		http.StatusUnsupportedMediaType:         UnsupportedMediaType,
		http.StatusRequestedRangeNotSatisfiable: RequestedRangeNotSatisfiable,
		http.StatusExpectationFailed:            ExpectationFailed,
		http.StatusUnprocessableEntity:          UnprocessableEntity,
		http.StatusTooManyRequests:              TooManyRequests,
		http.StatusInternalServerError:          ServerError,
		http.StatusNotImplemented:               NotImplemented,
		http.StatusBadGateway:                   BadGateway,
		http.StatusServiceUnavailable:           ServiceUnavailable,
		http.StatusGatewayTimeout:               GatewayTimeout,
	}

	for httpStatus, customCode := range statusMap {
		httpToCustomCode[httpStatus] = customCode
	}
}

// GetCustomCode maps an HTTP status to a custom response code.
func getCustomCode(httpStatus int) int {
	if code, exists := httpToCustomCode[httpStatus]; exists {
		return code
	}
	return ServerError // Default to internal server error if not mapped
}
