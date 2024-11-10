package middlewares

import (
	"bufio"
	"github.com/illusory-server/accounts/pkg/logger"
	"github.com/illusory-server/accounts/pkg/logger/log"
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
			log.String("http-scheme", scheme),
			log.String("http-proto", proto),
			log.String("http-method", method),
			log.String("remote-addr", remoteAddr),
			log.String("user-agent", userAgent),
			log.String("uri", uri),
			log.Duration("duration", time.Since(t)),
			log.Int("http-status", wRec.Status()),
		}

		if wRec.Status() > 400 {
			l.Error(ctx, "http error", fields...)
			return
		}
		l.Debug(ctx, "http ok", fields...)
	})
}
