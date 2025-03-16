package codes

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

type Code int

const (
	OK                 Code = 0
	Canceled           Code = 1
	Unknown            Code = 2
	InvalidArgument    Code = 3
	DeadlineExceeded   Code = 4
	NotFound           Code = 5
	AlreadyExists      Code = 6
	PermissionDenied   Code = 7
	ResourceExhausted  Code = 8
	FailedPrecondition Code = 9
	Aborted            Code = 10
	OutOfRange         Code = 11
	Unimplemented      Code = 12
	Internal           Code = 13
	Unavailable        Code = 14
	DataLoss           Code = 15
	Unauthenticated    Code = 16
)

func ToGRPC(code Code) codes.Code {
	return codes.Code(code)
}

func ToHTTP(code Code) int {
	switch code {
	case OK:
		return http.StatusOK
	case Canceled:
		return http.StatusRequestTimeout
	case Unknown:
		return http.StatusInternalServerError
	case InvalidArgument:
		return http.StatusBadRequest
	case DeadlineExceeded:
		return http.StatusRequestTimeout
	case NotFound:
		return http.StatusNotFound
	case AlreadyExists:
		return http.StatusConflict
	case PermissionDenied:
		return http.StatusForbidden
	case ResourceExhausted:
		return http.StatusTooManyRequests
	case FailedPrecondition:
		return http.StatusPreconditionFailed
	case Aborted:
		return http.StatusConflict
	case OutOfRange:
		return http.StatusBadRequest
	case Unimplemented:
		return http.StatusNotImplemented
	case Internal:
		return http.StatusInternalServerError
	case Unavailable:
		return http.StatusServiceUnavailable
	case DataLoss:
		return http.StatusInternalServerError
	case Unauthenticated:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

func FromGRPC(code codes.Code) Code {
	return Code(code)
}

func FromHTTP(httpCode int) Code {
	switch httpCode {
	case http.StatusOK:
		return OK
	case http.StatusRequestTimeout, http.StatusGatewayTimeout:
		return Canceled
	case http.StatusBadRequest:
		return InvalidArgument
	case http.StatusNotFound:
		return NotFound
	case http.StatusConflict:
		return AlreadyExists
	case http.StatusForbidden:
		return PermissionDenied
	case http.StatusTooManyRequests:
		return ResourceExhausted
	case http.StatusPreconditionFailed:
		return FailedPrecondition
	case http.StatusNotImplemented:
		return Unimplemented
	case http.StatusInternalServerError:
		return Internal
	case http.StatusServiceUnavailable:
		return Unavailable
	case http.StatusUnauthorized:
		return Unauthenticated
	default:
		if httpCode >= 400 && httpCode < 500 {
			return InvalidArgument
		}
		if httpCode >= 500 && httpCode < 600 {
			return Internal
		}
		return Unknown
	}
}
