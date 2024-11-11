package appecosystem

import (
	"context"
	"fmt"
	"github.com/illusory-server/accounts/pkg/app"
	"github.com/illusory-server/accounts/pkg/middlewares"
	"github.com/pkg/errors"
	"net/http"
	"sync"
	"time"
)

type (
	HTTPJobConfig struct {
		Address        string
		ReadTimeout    time.Duration
		WriteTimeout   time.Duration
		MaxHeaderBytes int
	}

	HTTPJob struct {
		app     *app.App
		server  *http.Server
		mu      sync.Mutex
		handler http.Handler
		stop    chan struct{}
		config  *HTTPJobConfig
	}
)

func (h *HTTPJob) Init(ctx context.Context, config any) error {
	return nil
}

func (h *HTTPJob) Run(ctx context.Context) error {
	if h.handler == nil {
		<-h.stop
	}

	handler := h.handler

	handler = middlewares.Sentry(handler)
	handler = middlewares.Logging(handler, h.app.Logger())
	handler = middlewares.Tracer(handler, h.app.Tracer())
	handler = middlewares.Prometheus(handler)

	s := &http.Server{
		Addr:           h.config.Address,
		Handler:        handler,
		ReadTimeout:    h.config.ReadTimeout,
		WriteTimeout:   h.config.WriteTimeout,
		MaxHeaderBytes: h.config.MaxHeaderBytes,
	}

	h.mu.Lock()
	h.server = s
	h.mu.Unlock()

	if err := h.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("http server is down: %w", err)
	}
	return nil
}

func (h *HTTPJob) Close(ctx context.Context) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	close(h.stop)

	if h.server != nil {
		return h.server.Shutdown(ctx)
	}

	return nil
}
