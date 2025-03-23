package middlewares

import (
	"bufio"
	"github.com/illusory-server/accounts/pkg/logger"
	"github.com/pkg/errors"
	"net"
	"net/http"
	"strings"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func newStatusRecorder(w http.ResponseWriter) *statusRecorder {
	return &statusRecorder{
		ResponseWriter: w,
	}
}

func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func (r *statusRecorder) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := r.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("hijack not supported")
	}
	return h.Hijack()
}

func (r *statusRecorder) Flush() {
	flusher, ok := r.ResponseWriter.(http.Flusher)
	if ok {
		flusher.Flush()
	}
}

func (r *statusRecorder) Status() int {
	if r.status == 0 {
		return http.StatusOK
	}
	return r.status
}

func Logging(handler http.Handler, l logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Liveness-Probe") == "Healthz" {
			handler.ServeHTTP(w, r)
			return
		}

		ctx := l.InjectCtx(r.Context())
		var scheme string
		if r.TLS != nil {
			scheme = "https"
		} else {
			scheme = "http"
		}

		proto := r.Proto
		method := r.Method
		remoteAddr := r.RemoteAddr
		userAgent := r.UserAgent()
		uri := strings.Join([]string{scheme, "://", r.Host, r.RequestURI}, "")
		wRec := newStatusRecorder(w)

		t := time.Now()

		handler.ServeHTTP(wRec, r)

		fields := []logger.Field{
			logger.String("http-scheme", scheme),
			logger.String("http-proto", proto),
			logger.String("http-method", method),
			logger.String("remote-addr", remoteAddr),
			logger.String("user-agent", userAgent),
			logger.String("uri", uri),
			logger.Duration("duration", time.Since(t)),
			logger.Int("http-status", wRec.Status()),
		}

		if wRec.Status() > 400 {
			l.Error(ctx, "http error", fields...)
			return
		}
		l.Debug(ctx, "http ok", fields...)
	})
}
