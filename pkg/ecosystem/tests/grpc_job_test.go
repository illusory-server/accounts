package grpc_job_test

import (
	"context"
	ayaka "github.com/illusory-server/accounts/pkg/core"
	"github.com/illusory-server/accounts/pkg/ecosystem"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func noopRegister(ctx context.Context, di ayaka.Container, srv *grpc.Server) error {
	return nil
}

func noopServerRegister(srv *grpc.Server) error {
	return nil
}

func noopInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	return nil, nil
}

func TestGrpcJobBuilder(t *testing.T) {
	t.Run("Should correct build grpc job", func(t *testing.T) {
		address := "localhost:10101"
		maxRetry := 5
		requestTimeout := time.Second * 5

		builder := ecosystem.NewGrpcJobBuilder()
		job, err := builder.
			Address(address).
			MaxRetry(maxRetry).
			RequestTimeout(requestTimeout).
			Register().
			RegisterOptions().
			RegisterServer().
			Interceptors().
			Build()

		assert.NoError(t, err)
		assert.NotNil(t, job)

		assert.Equal(t, address, job.Address())
		assert.Equal(t, maxRetry, job.MaxRetry())
		assert.Equal(t, requestTimeout, job.RequestTimeout())

		assert.Equal(t, 0, len(job.ServerRegs()))
		assert.Equal(t, 0, len(job.Regs()))
		assert.Equal(t, 0, len(job.Interceptors()))
		assert.Equal(t, 0, len(job.Options()))
	})

	t.Run("Should correct error building grpc without address, max-retry, request-timeout", func(t *testing.T) {
		address := "localhost:10101"
		maxRetry := 5

		builder := ecosystem.NewGrpcJobBuilder()
		job, err := builder.
			MaxRetry(maxRetry).
			RequestTimeout(time.Second * 5).
			Build()

		assert.Error(t, err)
		assert.Nil(t, job)

		builder = ecosystem.NewGrpcJobBuilder()
		job, err = builder.
			Address(address).
			RequestTimeout(time.Second * 5).
			Build()

		assert.Error(t, err)
		assert.Nil(t, job)

		builder = ecosystem.NewGrpcJobBuilder()
		job, err = builder.
			Address(address).
			MaxRetry(maxRetry).
			Build()

		assert.Error(t, err)
		assert.Nil(t, job)
	})

	t.Run("Should correct work register, serverRegister and interceptor", func(t *testing.T) {
		address := "localhost:10101"
		maxRetry := 5
		requestTimeout := time.Second * 5

		builder := ecosystem.NewGrpcJobBuilder()
		job, err := builder.
			Address(address).
			MaxRetry(maxRetry).
			RequestTimeout(requestTimeout).
			Register(noopRegister, noopRegister).
			RegisterServer(noopServerRegister, noopServerRegister, noopServerRegister).
			Interceptors(noopInterceptor, noopInterceptor, noopInterceptor, noopInterceptor).
			RegisterOptions(grpc.ChainUnaryInterceptor(), grpc.ChainUnaryInterceptor(), grpc.ChainUnaryInterceptor()).
			Build()

		assert.NoError(t, err)
		assert.NotNil(t, job)

		assert.Equal(t, address, job.Address())
		assert.Equal(t, maxRetry, job.MaxRetry())
		assert.Equal(t, requestTimeout, job.RequestTimeout())

		assert.Equal(t, 3, len(job.ServerRegs()))
		assert.Equal(t, 2, len(job.Regs()))
		assert.Equal(t, 4, len(job.Interceptors()))
		assert.Equal(t, 3, len(job.Options()))
	})
}

func TestGrpcJobSignature(t *testing.T) {
	address := "localhost:10101"
	maxRetry := 5
	requestTimeout := time.Second * 5

	builder := ecosystem.NewGrpcJobBuilder()
	job, err := builder.
		Address(address).
		MaxRetry(maxRetry).
		RequestTimeout(requestTimeout).
		Register().
		RegisterOptions().
		RegisterServer().
		Interceptors().
		Build()

	assert.NoError(t, err)

	ayaka.NewApp(&ayaka.Options{
		Name:        "aya",
		Description: "kekw",
		Version:     "0.0.1",
		Container:   ayaka.NewContainer(ayaka.NoopLogger{}),
	}).WithJob(ayaka.JobEntry{
		Key: "xd",
		Job: job,
	})
}
