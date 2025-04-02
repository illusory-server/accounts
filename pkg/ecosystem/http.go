package ecosystem

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation"
	ayaka "github.com/illusory-server/accounts/pkg/core"
	"github.com/pkg/errors"
)

const (
	DefaultHttpIdleTimeout         = 60 * time.Second
	DefaultHttpHeaderMaxBytesLimit = 4 * 1024 * 1024
)

type (
	HttpJobBuilder struct {
		address        string
		requestTimeout time.Duration
		idleTimeout    time.Duration
		maxHeaderBytes int

		regs        []HttpRegister
		middlewares []func(http.Handler) http.Handler
	}

	HttpJob struct {
		address        string
		requestTimeout time.Duration
		idleTimeout    time.Duration
		maxHeaderBytes int

		handler     *chi.Mux
		regs        []HttpRegister
		middlewares []func(http.Handler) http.Handler
	}

	HttpRegister func(ctx context.Context, di ayaka.Container, handler *chi.Mux) (*chi.Mux, error)
)

func (h *HttpJob) Address() string {
	return h.address
}

func (h *HttpJob) RequestTimeout() time.Duration {
	return h.requestTimeout
}

func (h *HttpJob) IdleTimeout() time.Duration {
	return h.idleTimeout
}

func (h *HttpJob) MaxHeaderBytes() int {
	return h.maxHeaderBytes
}

func (h *HttpJob) Handler() http.Handler {
	return h.handler
}

func (h *HttpJob) Regs() []HttpRegister {
	return h.regs
}

func (h *HttpJob) Middlewares() []func(http.Handler) http.Handler {
	return h.middlewares
}

func (h *HttpJob) Init(ctx context.Context, container ayaka.Container) error {
	errCh := make(chan error, 1)
	go func(errCh chan<- error) {
		var err error
		h.handler.Use(h.middlewares...)
		for _, reg := range h.regs {
			if h.handler, err = reg(ctx, container, h.handler); err != nil {
				errCh <- errors.Wrap(err, "[HttpJob] http register failed")
				return
			}
		}
		errCh <- nil
	}(errCh)

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		return nil
	}
}

func (h *HttpJob) Run(ctx context.Context, container ayaka.Container) error {
	errCh := make(chan error, 1)
	var log ayaka.Logger
	err := container.Invoke(func(logger ayaka.Logger) {
		log = logger
	})
	if err != nil {
		return errors.Wrap(err, "[GrpcJob] di.Invoke")
	}

	srv := http.Server{
		Addr:           h.address,
		Handler:        h.handler,
		WriteTimeout:   h.requestTimeout,
		ReadTimeout:    h.requestTimeout,
		IdleTimeout:    h.idleTimeout,
		MaxHeaderBytes: h.maxHeaderBytes,
	}

	go func() {
		log.Info(ctx, "http server started...", map[string]any{"address": h.address})
		err = srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
		errCh <- nil
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		err := srv.Shutdown(ctx)
		if err != nil {
			log.Warn(ctx, "http server failed graceful stopped", map[string]any{"address": h.address})
			return errors.Wrap(err, "[HttpJob] failed to shutdown http server")
		}
		log.Warn(ctx, "http server stopped", map[string]any{"address": h.address})
		return nil
	}
}

func NewHttpJobBuilder() *HttpJobBuilder {
	return &HttpJobBuilder{
		regs:        make([]HttpRegister, 0, sliceCap),
		middlewares: make([]func(http.Handler) http.Handler, 0, sliceCap),
	}
}

func (b *HttpJobBuilder) Address(address string) *HttpJobBuilder {
	b.address = address
	return b
}

func (b *HttpJobBuilder) RequestTimeout(requestTimeout time.Duration) *HttpJobBuilder {
	b.requestTimeout = requestTimeout
	return b
}

func (b *HttpJobBuilder) IdleTimeout(idleTimeout time.Duration) *HttpJobBuilder {
	b.idleTimeout = idleTimeout
	return b
}

func (b *HttpJobBuilder) MaxHeaderBytes(maxHeaderBytes int) *HttpJobBuilder {
	b.maxHeaderBytes = maxHeaderBytes
	return b
}

func (b *HttpJobBuilder) Middleware(middlewares ...func(http.Handler) http.Handler) *HttpJobBuilder {
	if len(middlewares) > 0 {
		b.middlewares = append(b.middlewares, middlewares...)
	}
	return b
}

func (b *HttpJobBuilder) Register(regs ...HttpRegister) *HttpJobBuilder {
	if len(regs) > 0 {
		b.regs = append(b.regs, regs...)
	}
	return b
}

func (b *HttpJobBuilder) Validate() error {
	return validation.ValidateStruct(b,
		validation.Field(&b.address, validation.Required),
		validation.Field(&b.requestTimeout, validation.Required),
	)
}

func (b *HttpJobBuilder) Build() (*HttpJob, error) {
	if err := b.Validate(); err != nil {
		return nil, errors.Wrap(err, "[HttpJob] validation failed")
	}
	if b.idleTimeout == 0 {
		b.idleTimeout = DefaultHttpIdleTimeout
	}
	if b.maxHeaderBytes == 0 {
		b.maxHeaderBytes = DefaultHttpHeaderMaxBytesLimit
	}

	return &HttpJob{
		address:        b.address,
		requestTimeout: b.requestTimeout,
		idleTimeout:    b.idleTimeout,
		maxHeaderBytes: b.maxHeaderBytes,
		handler:        chi.NewRouter(),
		regs:           b.regs,
		middlewares:    b.middlewares,
	}, nil
}
