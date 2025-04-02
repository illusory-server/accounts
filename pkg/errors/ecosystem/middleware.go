package ecosystem

import (
	"context"

	"github.com/illusory-server/accounts/pkg/errors/errx"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func UnaryErrorHandleInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		_ *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		result, err := handler(ctx, req)
		if err != nil {
			var libErr *errx.Error
			if errors.As(err, &libErr) {
				return nil, status.Error(ToGRPC(libErr.Code()), libErr.Error())
			}
			return nil, err
		}
		return result, nil
	}
}
