package middlewares

import (
	"github.com/opentracing/opentracing-go"
	"net/http"
)

func Tracer(handler http.Handler, tracer opentracing.Tracer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})
}
