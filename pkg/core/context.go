package ayaka

import (
	"context"
)

type appKey struct{}

func AppWithContext(ctx context.Context, app *App) context.Context {
	return context.WithValue(ctx, appKey{}, app)
}

func AppFromContext(ctx context.Context) (*ReadonlyApp, error) {
	val := ctx.Value(appKey{})
	if val == nil {
		return nil, ErrAppNotFountInContext
	}
	result := val.(*App)
	return &ReadonlyApp{
		app: result,
	}, nil
}
