package app

import (
	"context"
	"github.com/OddEer0/Eer0/app/eapp"
	"github.com/illusory-server/accounts/pkg/logger"
	"github.com/opentracing/opentracing-go"
	"io"
)

type Job interface {
	eapp.Job
}

type JobOpt struct {
	Key string
	Job Job
}

type Handler struct {
	Key     string
	Handler func(ctx context.Context, conf any) error
}

type App struct {
	app    *eapp.App
	env    *envVar
	err    error
	logger logger.Logger
	tracer opentracing.Tracer
	closer []io.Closer
}

func (a *App) Logger() logger.Logger {
	return a.logger
}

func (a *App) WithConfig(conf any) *App {
	if a.err != nil {
		return a
	}

	a.app = a.app.WithConfig(conf)
	return a
}

func (a *App) Config() any {
	return a.app.Configs().Client
}

func (a *App) Tracer() opentracing.Tracer {
	return a.tracer
}

func (a *App) Start() error {
	if a.err != nil {
		return a.err
	}
	err := a.app.Start()
	if err != nil {
		return err
	}
	return nil
}
