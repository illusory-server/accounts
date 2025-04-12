package main

import (
	"context"
	ayaka "github.com/OddEer0/ayaka/core"
	"github.com/OddEer0/ayaka/ecosystem"
	"github.com/illusory-server/accounts/cmd/dependency"
	v1 "github.com/illusory-server/accounts/gen/accounts/v1"
	grpcv1 "github.com/illusory-server/accounts/internal/transport/grpc/v1"
	"github.com/illusory-server/accounts/pkg/job"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"os"
)

func grpcServer(_ context.Context, container dependency.Dependency, srv *grpc.Server) error {
	server := grpcv1.NewServer(container.AccountUseCase())
	v1.RegisterAccountsServiceServer(srv, server)

	return nil
}

func main() {
	dependencyFactory := dependency.NewFactory()
	//ctx := context.Background()
	deps := dependencyFactory.Dependency()

	_ = os.Setenv("AYAKA_START_TIMEOUT", "5")
	_ = os.Setenv("AYAKA_GRACEFUL_TIMEOUT", "15")

	_ = os.Setenv("GRPC_ADDRESS", "localhost:11111")
	_ = os.Setenv("GRPC_REQUEST_TIMEOUT", "25")
	_ = os.Setenv("MON_ADDRESS", "localhost:11112")
	_ = os.Setenv("MON_TIMEOUT", "20")

	app := ayaka.NewApp(&ayaka.Options[dependency.Dependency]{
		Name:              "Accounts",
		Description:       "Core accounts service",
		Version:           "0.0.1",
		ConfigInterceptor: ecosystem.AdapterParseConfigFromEnv,
		Container:         deps,
		Logger:            ecosystem.NewAppLoggerWithZerolog(deps.Logger().Logger),
	}).
		WithConfig(&ayaka.Config{}).
		WithJob(
			ayaka.JobEntry[dependency.Dependency]{
				Key: "grpc-server",
				Job: job.MustGrpcJobEnv(
					nil,
					opentracing.NoopTracer{},
					deps.Logger(),
					grpcServer,
				),
			},
			ayaka.JobEntry[dependency.Dependency]{
				Key: "mon-server",
				Job: job.MustMonJobEnv[dependency.Dependency](nil),
			},
		)

	err := ecosystem.StartWithCli(app, ecosystem.DefaultPrinter{})
	if err != nil {
		panic(err)
	}
}
