package ecosystem

import (
	"context"
	"net"
	"sync"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	ayaka "github.com/illusory-server/accounts/pkg/core"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	sliceCap = 8
)

type (
	GrpcJobBuilder struct {
		address        string
		requestTimeout time.Duration
		interceptors   []grpc.UnaryServerInterceptor
		regs           []GrpcRegister
		serverRegs     []GrpcServerRegister
		options        []grpc.ServerOption
	}

	GrpcJob struct {
		srv            *grpc.Server
		mu             sync.Mutex
		address        string
		requestTimeout time.Duration
		interceptors   []grpc.UnaryServerInterceptor
		regs           []GrpcRegister
		serverRegs     []GrpcServerRegister
		options        []grpc.ServerOption
	}

	GrpcRegister       func(ctx context.Context, di ayaka.Container, srv *grpc.Server) error
	GrpcServerRegister func(srv *grpc.Server) error
)

func (g *GrpcJob) Address() string {
	return g.address
}

func (g *GrpcJob) RequestTimeout() time.Duration {
	return g.requestTimeout
}

func (g *GrpcJob) Interceptors() []grpc.UnaryServerInterceptor {
	return g.interceptors
}

func (g *GrpcJob) Regs() []GrpcRegister {
	return g.regs
}

func (g *GrpcJob) ServerRegs() []GrpcServerRegister {
	return g.serverRegs
}

func (g *GrpcJob) Options() []grpc.ServerOption {
	return g.options
}

func (g *GrpcJobBuilder) Validate() error {
	return validation.ValidateStruct(g,
		validation.Field(&g.address, validation.Required),
		validation.Field(&g.requestTimeout, validation.Required),
	)
}

func NewGrpcJobBuilder() *GrpcJobBuilder {
	return &GrpcJobBuilder{
		regs:         make([]GrpcRegister, 0, sliceCap),
		serverRegs:   make([]GrpcServerRegister, 0, sliceCap),
		interceptors: make([]grpc.UnaryServerInterceptor, 0, sliceCap),
		options:      make([]grpc.ServerOption, 0, sliceCap),
	}
}

func (g *GrpcJob) Init(ctx context.Context, di ayaka.Container) error {
	sliceInterceptors := make([]grpc.UnaryServerInterceptor, 0, len(g.interceptors))
	copy(sliceInterceptors, g.interceptors)

	if g.requestTimeout > 0 {
		sliceInterceptors = append(sliceInterceptors, TimeoutInterceptor(g.requestTimeout))
	}

	grpcOptions := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(sliceInterceptors...),
	}
	grpcOptions = append(grpcOptions, g.options...)

	srv := grpc.NewServer(grpcOptions...)

	errCh := make(chan error, 1)
	go func(errCh chan<- error) {
		for _, reg := range g.regs {
			if err := reg(ctx, di, srv); err != nil {
				errCh <- errors.Wrap(err, "[GrpcJob] grpc register error")
				return
			}
		}

		for _, serverRegister := range g.serverRegs {
			if err := serverRegister(srv); err != nil {
				errCh <- errors.Wrap(err, "[GrpcJob] grpc register error")
				return
			}
		}
		errCh <- nil
	}(errCh)

	g.mu.Lock()
	g.srv = srv
	g.mu.Unlock()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		return nil
	}
}

func (g *GrpcJob) Run(ctx context.Context, di ayaka.Container) error {
	errCh := make(chan error, 1)
	var log ayaka.Logger
	err := di.Invoke(func(logger ayaka.Logger) {
		log = logger
	})
	if err != nil {
		return errors.Wrap(err, "[GrpcJob] di.Invoke")
	}

	go func() {
		if g.srv != nil {
			log.Info(ctx, "grpc server started...", map[string]any{"address": g.address})

			lis, err := net.Listen("tcp", g.address)
			if err != nil {
				errCh <- errors.Wrap(err, "[GrpcJob] net.Listen")
				return
			}

			err = g.srv.Serve(lis)
			if err != nil {
				errCh <- errors.Wrap(err, "[GrpcJob] srv.Serve")
				return
			}
		}

		errCh <- nil
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		log.Warn(ctx, "grpc server stopped", map[string]any{"address": g.address})
		g.srv.GracefulStop()
		return nil
	}
}

func (g *GrpcJobBuilder) Address(address string) *GrpcJobBuilder {
	g.address = address
	return g
}

func (g *GrpcJobBuilder) RequestTimeout(timeout time.Duration) *GrpcJobBuilder {
	g.requestTimeout = timeout
	return g
}

func (g *GrpcJobBuilder) Interceptors(interceptors ...grpc.UnaryServerInterceptor) *GrpcJobBuilder {
	if len(interceptors) > 0 {
		g.interceptors = append(g.interceptors, interceptors...)
	}
	return g
}

func (g *GrpcJobBuilder) Register(regs ...GrpcRegister) *GrpcJobBuilder {
	if len(regs) > 0 {
		g.regs = append(g.regs, regs...)
	}
	return g
}

func (g *GrpcJobBuilder) RegisterServer(regs ...GrpcServerRegister) *GrpcJobBuilder {
	if len(regs) > 0 {
		g.serverRegs = append(g.serverRegs, regs...)
	}
	return g
}

func (g *GrpcJobBuilder) RegisterOptions(options ...grpc.ServerOption) *GrpcJobBuilder {
	if len(options) > 0 {
		g.options = append(g.options, options...)
	}
	return g
}

func (g *GrpcJobBuilder) Build() (*GrpcJob, error) {
	err := g.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "[GrpcJobBuilder] validate error")
	}

	return &GrpcJob{
		address:        g.address,
		requestTimeout: g.requestTimeout,
		interceptors:   g.interceptors,
		regs:           g.regs,
		serverRegs:     g.serverRegs,
		options:        g.options,
		mu:             sync.Mutex{},
	}, nil
}
