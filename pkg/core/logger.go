package ayaka

import "context"

type Logger interface {
	Debug(ctx context.Context, message string, info map[string]any)
	Info(ctx context.Context, message string, info map[string]any)
	Warn(ctx context.Context, message string, info map[string]any)
	Error(ctx context.Context, message string, info map[string]any)
}

type noopLogger struct{}

func (n noopLogger) Debug(ctx context.Context, message string, info map[string]any) {}

func (n noopLogger) Info(ctx context.Context, message string, info map[string]any) {}

func (n noopLogger) Warn(ctx context.Context, message string, info map[string]any) {}

func (n noopLogger) Error(ctx context.Context, message string, info map[string]any) {}
