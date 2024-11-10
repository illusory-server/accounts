package app

import (
	"context"
	"github.com/pkg/errors"
)

type ContextKey struct{}

var ErrAppCannotFromContext = errors.New("app cannot from context")

func FromContext(ctx context.Context) (ReadApp, error) {
	app, ok := ctx.Value(ContextKey{}).(*App)
	if !ok {
		return ReadApp{}, ErrAppCannotFromContext
	}
	return ReadApp{
		app: app,
	}, nil
}
