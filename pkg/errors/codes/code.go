package codes

type Code int

const (
	NotCode            Code = 0
	Ok                 Code = 200
	Created            Code = 201
	Accepted           Code = 202
	BadRequest         Code = 400
	Unauthorized       Code = 401
	Forbidden          Code = 403
	NotFound           Code = 404
	DeadlineExceeded   Code = 408
	Conflict           Code = 409
	Unprocessable      Code = 422
	Internal           Code = 500
	Unimplemented      Code = 501
	BadGateway         Code = 502
	ServiceUnavailable Code = 503
	GatewayTimeout     Code = 504
)
