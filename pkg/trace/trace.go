package trace

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

func ExtractTraceID(ctx context.Context) string {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return ""
	}

	return ExtractTraceFromSpan(span)
}

func ExtractTraceFromSpan(span opentracing.Span) string {
	jaegerSpanContext, ok := span.Context().(jaeger.SpanContext)
	if !ok {
		return ""
	}

	return jaegerSpanContext.TraceID().String()
}
