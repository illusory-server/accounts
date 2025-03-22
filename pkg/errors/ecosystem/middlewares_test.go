package ecosystem

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	libCodes "github.com/illusory-server/accounts/pkg/errors/codex"
	libErr "github.com/illusory-server/accounts/pkg/errors/errx"
)

// nolint:testifylint
func TestUnaryErrorHandleInterceptor(t *testing.T) {
	tests := []struct {
		name        string
		handlerErr  error
		expectedErr error
	}{
		{
			name:        "no error",
			handlerErr:  nil,
			expectedErr: nil,
		},
		{
			name:        "custom error 1",
			handlerErr:  libErr.New(libCodes.InvalidArgument, "invalid argument"),
			expectedErr: status.Error(codes.InvalidArgument, "invalid argument"),
		},
		{
			name:        "custom error 2",
			handlerErr:  libErr.New(libCodes.Unauthenticated, "authorization failed"),
			expectedErr: status.Error(codes.Unauthenticated, "authorization failed"),
		},
		{
			name:        "generic error",
			handlerErr:  errors.New("generic error"),
			expectedErr: errors.New("generic error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interceptor := UnaryErrorHandleInterceptor()

			handler := func(_ context.Context, _ interface{}) (interface{}, error) {
				return nil, tt.handlerErr
			}

			_, err := interceptor(context.Background(), nil, nil, handler)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), tt.expectedErr.Error())
			}
		})
	}
}

func TestMultiServiceErrorHandler(t *testing.T) {
	interceptor := UnaryErrorHandleInterceptor()

	causeMsg := "not found"
	errOrig := libErr.New(libCodes.NotFound, causeMsg)
	var err error
	err = errors.Wrap(errOrig, "wrap1")

	_, errFromInterceptor := interceptor(context.Background(), nil, nil,
		func(_ context.Context, _ interface{}) (interface{}, error) {
			return nil, err
		},
	)
	grpcStatusFromClient, ok := status.FromError(errFromInterceptor)
	assert.True(t, ok)
	assert.Equal(t, codes.NotFound, grpcStatusFromClient.Code())

	err = errors.Wrap(grpcStatusFromClient.Err(), "wrap2")
	err = errors.Wrap(grpcStatusFromClient.Err(), "wrap3")

	_, errFromInterceptor = interceptor(context.Background(), nil, nil,
		func(_ context.Context, _ interface{}) (interface{}, error) {
			return nil, err
		},
	)

	grpcStatusFromClient, ok = status.FromError(errFromInterceptor)
	assert.True(t, ok)
	assert.Equal(t, codes.NotFound, grpcStatusFromClient.Code())

	err = libErr.WrapWithCode(grpcStatusFromClient.Err(), libCodes.InvalidArgument, "wrap4")
	err = errors.Wrap(err, "wrap5")

	_, errFromInterceptor = interceptor(context.Background(), nil, nil,
		func(_ context.Context, _ interface{}) (interface{}, error) {
			return nil, err
		},
	)

	grpcStatusFromClient, ok = status.FromError(errFromInterceptor)
	assert.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, grpcStatusFromClient.Code())
}
