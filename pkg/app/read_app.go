package app

import (
	"github.com/illusory-server/accounts/pkg/logger"
	"github.com/opentracing/opentracing-go"
)

type ReadApp struct {
	app *App
}

func (r ReadApp) Config() any {
	return r.app.Config()
}

func (r ReadApp) Tracer() opentracing.Tracer {
	return r.app.Tracer()
}

func (r ReadApp) Logger() logger.Logger {
	return r.app.Logger()
}
