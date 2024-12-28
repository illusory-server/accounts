package job

import (
	"github.com/illusory-server/accounts/pkg/ecosystem"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"os"
	"strconv"
	"time"
)

type GrpcEnvJobEnvKeys struct {
	Address,
	RequestTimeout string
}

var (
	DefaultUnaryJobEnvKeys = GrpcEnvJobEnvKeys{
		Address:        "GRPC_ADDRESS",
		RequestTimeout: "GRPC_REQUEST_TIMEOUT",
	}
)

// MustGrpcEnvJob default key - GRPC_ADDRESS, GRPC_REQUEST_TIMEOUT, GRPC_MAX_RETRY
func MustGrpcEnvJob(
	keys *GrpcEnvJobEnvKeys,
	tracer opentracing.Tracer,
	regs ...ecosystem.GrpcRegister,
) *ecosystem.GrpcJob {
	if keys == nil {
		keys = &DefaultUnaryJobEnvKeys
	}

	address := os.Getenv(keys.Address)
	if address == "" {
		panic("environment variable '" + keys.Address + "' is not set")
	}
	requestTimeoutEnv := os.Getenv(keys.RequestTimeout)
	if requestTimeoutEnv == "" {
		panic("environment variable '" + keys.RequestTimeout + "' is not set")
	}
	requestTimeout, err := strconv.Atoi(requestTimeoutEnv)
	if err != nil {
		panic("environment variable '" + keys.RequestTimeout + "' is not a number")
	}

	job, err := ecosystem.NewGrpcJobBuilder().
		Address(address).
		RequestTimeout(time.Duration(requestTimeout) * time.Second).
		Register(regs...).
		RegisterServer(func(srv *grpc.Server) error {
			// Register monitoring
			registerPrometheus(srv)
			// Register healthcheck service
			health.RegisterHealthServer(srv, new(healthService))
			// Register reflection service on gRPC server.
			reflection.Register(srv)
			return nil
		}).
		Build()

	if err != nil {
		panic(err)
	}

	return job
}
