package job

import (
	"github.com/OddEer0/ayaka/ecosystem"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/illusory-server/accounts/pkg/interceptors"
	"github.com/illusory-server/accounts/pkg/logger"
	"github.com/illusory-server/accounts/pkg/safe"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"time"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type (
	GrpcJobParams struct {
		Address        string
		RequestTimeout time.Duration
	}
	GrpcEnvJobEnvKeys struct {
		Address,
		RequestTimeout string
	}
)

var (
	DefaultUnaryJobEnvKeys = GrpcEnvJobEnvKeys{
		Address:        "GRPC_ADDRESS",
		RequestTimeout: "GRPC_REQUEST_TIMEOUT",
	}
)

func getGrpcEnvJob(keys *GrpcEnvJobEnvKeys) (GrpcJobParams, error) {
	if keys == nil {
		keys = &DefaultUnaryJobEnvKeys
	}

	address := os.Getenv(keys.Address)
	if address == "" {
		return GrpcJobParams{}, errors.Errorf("environment variable %s is not set", keys.Address)
	}
	requestTimeoutEnv := os.Getenv(keys.RequestTimeout)
	if requestTimeoutEnv == "" {
		return GrpcJobParams{}, errors.Errorf("environment variable %s is not set", keys.RequestTimeout)
	}
	requestTimeout, err := strconv.Atoi(requestTimeoutEnv)
	if err != nil {
		return GrpcJobParams{}, errors.Errorf("environment variable %s is not a number", keys.RequestTimeout)
	}

	return GrpcJobParams{
		Address:        address,
		RequestTimeout: time.Duration(requestTimeout) * time.Second,
	}, nil
}

// NewGrpcJob default key - GRPC_ADDRESS, GRPC_REQUEST_TIMEOUT, GRPC_MAX_RETRY
func NewGrpcJob[T any](
	params GrpcJobParams,
	tracer opentracing.Tracer,
	logger logger.Logger,
	regs ...ecosystem.GrpcRegister[T],
) (*ecosystem.GrpcJob[T], error) {
	job, err := ecosystem.NewGrpcJobBuilder[T]().
		Address(params.Address).
		RequestTimeout(params.RequestTimeout).
		RecoverHandler(safe.Recover).
		Interceptors(
			grpcPrometheus.UnaryServerInterceptor,
			otgrpc.OpenTracingServerInterceptor(tracer),
			interceptors.Logging(logger),
			grpcValidator.UnaryServerInterceptor(),
			interceptors.Sentry(),
		).
		RegisterOptions(
			grpc.ChainStreamInterceptor(
				grpcPrometheus.StreamServerInterceptor,
				otgrpc.OpenTracingStreamServerInterceptor(tracer),
				recovery.StreamServerInterceptor(),
				grpcPrometheus.StreamServerInterceptor,
				grpcValidator.StreamServerInterceptor(),
			),
		).
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
		return nil, errors.Wrap(err, "ecosystem.NewGrpcJobBuilder.Build")
	}

	return job, nil
}

func NewGrpcJobEnv[T any](
	keys *GrpcEnvJobEnvKeys,
	tracer opentracing.Tracer,
	logger logger.Logger,
	regs ...ecosystem.GrpcRegister[T],
) (*ecosystem.GrpcJob[T], error) {
	params, err := getGrpcEnvJob(keys)
	if err != nil {
		return nil, errors.Wrap(err, "getGrpcEnvJob")
	}
	return NewGrpcJob(params, tracer, logger, regs...)
}

func MustGrpcJobEnv[T any](
	keys *GrpcEnvJobEnvKeys,
	tracer opentracing.Tracer,
	logger logger.Logger,
	regs ...ecosystem.GrpcRegister[T],
) *ecosystem.GrpcJob[T] {
	job, err := NewGrpcJobEnv(keys, tracer, logger, regs...)
	if err != nil {
		panic(err)
	}
	return job
}
