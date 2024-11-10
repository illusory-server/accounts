package app

import (
	"context"
	"github.com/OddEer0/Eer0/app/eapp"
	"github.com/OddEer0/Eer0/app/ecosystem"
	"github.com/illusory-server/accounts/pkg/logger/log"
)

type Options struct {
	Name, Version, Description string
}

func Init(opt *Options) *App {
	app := &App{}
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

	// TODO - Сделать инициализацию Tracer`а

	app.app = app.app.LibConfig(&eapp.LibConfig{
		StartTimeout: env.StartTimeout,
		StopTimeout:  env.StopTimeout,
	})

	return app
}
