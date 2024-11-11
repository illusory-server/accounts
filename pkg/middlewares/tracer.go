package middlewares

import (
	"github.com/illusory-server/accounts/pkg/trace"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

func Tracer(handler http.Handler, tracer opentracing.Tracer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		spanCtx, _ := tracer.Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))

		span, ctx := opentracing.StartSpanFromContextWithTracer(
			r.Context(), tracer, r.Method+" "+r.URL.Path,
			ext.RPCServerOption(spanCtx))
		defer span.Finish()

		ext.SpanKindRPCClient.Set(span)
		ext.HTTPUrl.Set(span, r.URL.String())
		ext.HTTPMethod.Set(span, r.Method)

		w.Header().Add("X-Trace-ID", trace.ExtractTraceFromSpan(span))

		wRec := newStatusRecorder(w)
		handler.ServeHTTP(wRec, r.WithContext(ctx))

		ext.HTTPStatusCode.Set(span, uint16(wRec.Status()))
		if wRec.Status() >= http.StatusBadRequest {
			ext.Error.Set(span, true)
		}
	})
}
