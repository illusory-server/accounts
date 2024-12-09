package xerr

import "github.com/illusory-server/accounts/pkg/errors/codes"

var (
	ErrNotFound           = New(codes.NotFound, "not found")
	ErrUnprocessable      = New(codes.Unprocessable, "unprocessable entity")
	ErrBadRequest         = New(codes.BadRequest, "bad request")
	ErrInternal           = New(codes.Internal, "internal error")
	ErrUnauthorized       = New(codes.Unauthorized, "unauthorized")
	ErrForbidden          = New(codes.Forbidden, "forbidden")
	ErrDeadlineExceeded   = New(codes.DeadlineExceeded, "deadline exceeded")
	ErrConflict           = New(codes.Conflict, "conflict")
	ErrUnimplemented      = New(codes.Unimplemented, "unimplemented")
	ErrBadGateway         = New(codes.BadGateway, "bad gateway")
	ErrServiceUnavailable = New(codes.ServiceUnavailable, "service unavailable")
	ErrGatewayTimeout     = New(codes.GatewayTimeout, "gateway timeout")
)
