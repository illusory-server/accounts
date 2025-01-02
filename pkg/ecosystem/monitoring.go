package ecosystem

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation"
	ayaka "github.com/illusory-server/accounts/pkg/core"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"net/http/pprof"
	"time"
)

const (
	DefaultMonitoringMaxHeaderBytes = 1 << 20
	DefaultMonitoringRequestTimeout = time.Minute
)

type (
	MonitoringJobBuilder struct {
		address string
		mux     *http.ServeMux
	}

	MonitoringJob struct {
		address string
		mux     *http.ServeMux
	}
)

func (m *MonitoringJobBuilder) Validate() error {
	return validation.ValidateStruct(m,
		validation.Field(&m.address, validation.Required),
	)
}

func (m *MonitoringJobBuilder) Build() (*MonitoringJob, error) {
	if err := m.Validate(); err != nil {
		return nil, err
	}
	if m.mux == nil {
		m.mux = http.NewServeMux()
	}
	return &MonitoringJob{
		address: m.address,
		mux:     m.mux,
	}, nil
}

func (m *MonitoringJob) Init(ctx context.Context, _ ayaka.Container) error {
	app, err := ayaka.AppFromContext(ctx)
	if err != nil {
		return errors.Wrap(err, "[MonitoringJob.Init] ayaka.AppFromContext")
	}

	// Register Prometheus metrics handler.
	m.mux.Handle("/metrics", promhttp.Handler())

	// Register pprof handlers
	m.mux.HandleFunc("/debug/pprof/", pprof.Index)
	m.mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	m.mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	m.mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	m.mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	err = prometheus.Register(prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Namespace:   "App",
			Name:        "info",
			Help:        "Application info",
			ConstLabels: prometheus.Labels{"version": app.Info().Version},
		},
		func() float64 { return float64(1) },
	))
	if err != nil {
		return errors.Wrap(err, "[MonitoringJob] prometheus.Register")
	}

	return nil
}

func (m *MonitoringJob) Run(ctx context.Context, container ayaka.Container) error {
	errCh := make(chan error, 1)

	var log ayaka.Logger
	err := container.Invoke(func(logger ayaka.Logger) {
		log = logger
	})
	if err != nil {
		return errors.Wrap(err, "[GrpcJob] di.Invoke")
	}

	srv := http.Server{
		Addr:           m.address,
		Handler:        m.mux,
		WriteTimeout:   DefaultMonitoringRequestTimeout,
		ReadTimeout:    DefaultMonitoringRequestTimeout,
		MaxHeaderBytes: DefaultMonitoringMaxHeaderBytes,
	}

	go func() {
		log.Info(ctx, "http monitoring server started...", map[string]any{"address": m.address})
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
			log.Warn(ctx, "http monitoring server failed graceful stopped", map[string]any{"address": m.address})
			return errors.Wrap(err, "[MonitoringJob] failed to shutdown http monitoring server")
		}
		log.Warn(ctx, "http server stopped", map[string]any{"address": m.address})
		return nil
	}
}

func (m *MonitoringJob) Address() string {
	return m.address
}

func (m *MonitoringJobBuilder) Address(addr string) *MonitoringJobBuilder {
	m.address = addr
	return m
}

func (m *MonitoringJobBuilder) Mux(mux *http.ServeMux) *MonitoringJobBuilder {
	m.mux = mux
	return m
}

func NewMonitoringJobBuilder() *MonitoringJobBuilder {
	return &MonitoringJobBuilder{}
}
