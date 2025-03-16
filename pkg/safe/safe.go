package safe

import (
	"github.com/getsentry/sentry-go"
	"time"
)

const flushTime = time.Second * 2

type RecoverOption interface {
	OnPanic(r any)
}

func Recover(opts ...RecoverOption) {
	if r := recover(); r != nil {
		hub := sentry.CurrentHub().Clone()
		hub.Recover(r)
		hub.Flush(flushTime)
		for _, opt := range opts {
			opt.OnPanic(r)
		}
	}
}

func Go(fn func()) {
	go func() {
		defer Recover()
		fn()
	}()
}
