package ecosystem

import (
	libCodes "github.com/illusory-server/accounts/pkg/errors/codes"
	"google.golang.org/grpc/codes"
	"net/http"
)

func ToGRPC(code libCodes.Code) codes.Code {
	return codes.Code(code)
}

func ToHTTP(code libCodes.Code) int {
	switch code {
	case libCodes.OK:
		return http.StatusOK
	case libCodes.Canceled:
		return http.StatusRequestTimeout
	case libCodes.Unknown:
		return http.StatusInternalServerError
	case libCodes.InvalidArgument:
		return http.StatusBadRequest
	case libCodes.DeadlineExceeded:
		return http.StatusRequestTimeout
	case libCodes.NotFound:
		return http.StatusNotFound
	case libCodes.AlreadyExists:
		return http.StatusConflict
	case libCodes.PermissionDenied:
		return http.StatusForbidden
	case libCodes.ResourceExhausted:
		return http.StatusTooManyRequests
	case libCodes.FailedPrecondition:
		return http.StatusPreconditionFailed
	case libCodes.Aborted:
		return http.StatusConflict
	case libCodes.OutOfRange:
		return http.StatusBadRequest
	case libCodes.Unimplemented:
		return http.StatusNotImplemented
	case libCodes.Internal:
		return http.StatusInternalServerError
	case libCodes.Unavailable:
		return http.StatusServiceUnavailable
	case libCodes.DataLoss:
		return http.StatusInternalServerError
	case libCodes.Unauthenticated:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

func FromGRPC(code codes.Code) libCodes.Code {
	return libCodes.Code(code)
}

func FromHTTP(httpCode int) libCodes.Code {
	switch httpCode {
	case http.StatusOK:
		return libCodes.OK
	case http.StatusRequestTimeout, http.StatusGatewayTimeout:
		return libCodes.Canceled
	case http.StatusBadRequest:
		return libCodes.InvalidArgument
	case http.StatusNotFound:
		return libCodes.NotFound
	case http.StatusConflict:
		return libCodes.AlreadyExists
	case http.StatusForbidden:
		return libCodes.PermissionDenied
	case http.StatusTooManyRequests:
		return libCodes.ResourceExhausted
	case http.StatusPreconditionFailed:
		return libCodes.FailedPrecondition
	case http.StatusNotImplemented:
		return libCodes.Unimplemented
	case http.StatusInternalServerError:
		return libCodes.Internal
	case http.StatusServiceUnavailable:
		return libCodes.Unavailable
	case http.StatusUnauthorized:
		return libCodes.Unauthenticated
	default:
		if httpCode >= 400 && httpCode < 500 {
			return libCodes.InvalidArgument
		}
		if httpCode >= 500 && httpCode < 600 {
			return libCodes.Internal
		}
		return libCodes.Unknown
	}
}
