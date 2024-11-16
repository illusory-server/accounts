package appecosystem

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	grpcvalidator "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/illusory-server/accounts/pkg/app"
	"github.com/illusory-server/accounts/pkg/interceptors"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"sync"
	"time"
)

type (
	GrpcJobConfig struct {
		Address        string
		RequestTimeout time.Duration
		MaxRetry       int
	}

	GrpcJob struct {
		config *GrpcJobConfig
		srv    *grpc.Server
		mu     sync.Mutex
		stop   chan struct{}
		regs   []GrpcRegister
		app    *app.App
	}

	GrpcRegister func(ctx context.Context, config interface{}, srv *grpc.Server) error
)

func (g *GrpcJob) Init(ctx context.Context, config any) error {
	tracer := g.app.Tracer()

	sliceInterceptors := []grpc.UnaryServerInterceptor{
		grpcprometheus.UnaryServerInterceptor,
		otgrpc.OpenTracingServerInterceptor(tracer),
		interceptors.Logging(g.app.Logger()),
		grpcvalidator.UnaryServerInterceptor(),
	}

	if g.config.RequestTimeout != 0 {
		sliceInterceptors = append(sliceInterceptors, interceptors.Timeout(g.config.RequestTimeout))
	}

	sliceInterceptors = append(sliceInterceptors, interceptors.Sentry())

	grpcOptions := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(sliceInterceptors...),
		grpc.ChainStreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(tracer),
			recovery.StreamServerInterceptor(),
			grpcprometheus.StreamServerInterceptor,
			grpcvalidator.StreamServerInterceptor(),
		),
	}

	srv := grpc.NewServer(grpcOptions...)

	for _, reg := range g.regs {
		if err := reg(ctx, config, srv); err != nil {
			return errors.Wrap(err, "[GrpcJob] grpc register error")
		}
	}

	// Register monitoring
	interceptors.RegisterPrometheus(srv)
	//// Register healthcheck service
	//health.RegisterHealthServer(srv, new(healthService))
	// Register reflection service on gRPC server.
	reflection.Register(srv)

	g.mu.Lock()
	g.srv = srv
	g.mu.Unlock()

	return nil
}

func (g *GrpcJob) Run(ctx context.Context) error {
	if g.srv == nil {
		<-g.stop
		return nil
	}

	cfg := g.config
	lis, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return errors.Wrap(err, "[GrpcJob] net.Listen")
	}
	err = g.srv.Serve(lis)
	if err != nil {
		return errors.Wrap(err, "[GrpcJob] srv.Serve")
	}

	return nil
}

func (g *GrpcJob) Close(ctx context.Context) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	close(g.stop)

	if g.srv != nil {
		stopped := make(chan struct{})
		go func() {
			g.srv.GracefulStop()
			close(stopped)
		}()

		// TODO - выяснить, нужно ли логировать/выводить ошибку при ctx.Done так как тут не получилось сделать GracefulStop
		select {
		case <-ctx.Done():
			g.srv.Stop()
		case <-stopped:
		}
	}

	return nil
}

func NewGrpcJob(app *app.App, conf *GrpcJobConfig, regs ...GrpcRegister) *GrpcJob {
	return &GrpcJob{
		config: conf,
		mu:     sync.Mutex{},
		stop:   make(chan struct{}),
		regs:   regs,
		app:    app,
	}
}
