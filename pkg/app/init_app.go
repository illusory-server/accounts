package app

import (
	"context"
	"github.com/OddEer0/Eer0/app/eapp"
	"github.com/OddEer0/Eer0/app/ecosystem"
	"github.com/illusory-server/accounts/pkg/logger/log"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics/prometheus"
	"io"
)

type Options struct {
	Name, Version, Description string
}

func Init(opt *Options) *App {
	app := &App{}
	app.closer = make([]io.Closer, 0, 8)
	env, err := parseEnv()
	if err != nil {
		app.err = err
		return app
	}
	app.app = eapp.Init(&eapp.InitOptions{
		Name:               opt.Name,
		Version:            opt.Version,
		Description:        opt.Description,
		UserCfgInterceptor: ecosystem.AdapterParseConfigFromYaml(env.ConfigPath),
	})

	app.env = env

	ctx := context.Background()
	app.logger = log.NewLogger(&log.Options{
		Level:  convertLogLvlToIntLvl(env.LogLevel),
		Pretty: env.LogPretty,
	})
	ctx = app.logger.InjectCtx(ctx)

	log.Info(ctx, "env parsed", log.Any("environment", env))

	// TODO - Сделать инициализацию Sentry

	err = app.initTracer(opt.Name)
	if err != nil {
		app.err = err
	}

	app.app = app.app.LibConfig(&eapp.LibConfig{
		StartTimeout: env.StartTimeout,
		StopTimeout:  env.StopTimeout,
	})

	return app
}

func (a *App) initTracer(name string) error {
	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		return errors.Wrap(err, "[App] jaegercfg.FromEnv")
	}

	cfg.ServiceName = name

	metricsFactory := prometheus.New()

	tracer, closer, err := cfg.NewTracer(jaegercfg.Metrics(metricsFactory))
	if err != nil {
		return errors.Wrap(err, "[App] cfg.NewTracer")
	}

	a.tracer = tracer
	a.closer = append(a.closer, closer)

	opentracing.SetGlobalTracer(a.tracer)

	return nil
}
