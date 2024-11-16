package interceptors

import (
	"context"
	"google.golang.org/grpc"
)

func Sentry() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		return handler(ctx, req)
	}
}
