package grpc_job_test

import (
	"github.com/illusory-server/accounts/pkg/ecosystem"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGrpcJobBuilder(t *testing.T) {
	t.Run("Should correct build grpc job", func(t *testing.T) {
		address := "localhost:10101"
		maxRetry := 5

		builder := ecosystem.NewGrpcJobBuilder()
		job, err := builder.
			Tracer(opentracing.NoopTracer{}).
			Address(address).
			MaxRetry(maxRetry).
			RequestTimeout(time.Second * 5).
			Build()

		assert.NoError(t, err)
		assert.NotNil(t, job)
	})

	t.Run("Should correct error building grpc without address, max-retry, tracer, request-timeout", func(t *testing.T) {
		address := "localhost:10101"
		maxRetry := 5

		builder := ecosystem.NewGrpcJobBuilder()
		job, err := builder.
			Address(address).
			MaxRetry(maxRetry).
			RequestTimeout(time.Second * 5).
			Build()

		assert.Error(t, err)
		assert.Nil(t, job)

		builder = ecosystem.NewGrpcJobBuilder()
		job, err = builder.
			Tracer(opentracing.NoopTracer{}).
			MaxRetry(maxRetry).
			RequestTimeout(time.Second * 5).
			Build()

		assert.Error(t, err)
		assert.Nil(t, job)

		builder = ecosystem.NewGrpcJobBuilder()
		job, err = builder.
			Tracer(opentracing.NoopTracer{}).
			Address(address).
			RequestTimeout(time.Second * 5).
			Build()

		assert.Error(t, err)
		assert.Nil(t, job)

		builder = ecosystem.NewGrpcJobBuilder()
		job, err = builder.
			Tracer(opentracing.NoopTracer{}).
			Address(address).
			MaxRetry(maxRetry).
			Build()

		assert.Error(t, err)
		assert.Nil(t, job)
	})
}
