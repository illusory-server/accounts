package safe

import (
	"time"

	"github.com/getsentry/sentry-go"
)

const flushTime = time.Second * 2

type RecoverOption interface {
	OnPanic(r any)
}

func RecoverWithOption(opts ...RecoverOption) {
	if r := recover(); r != nil {
		hub := sentry.CurrentHub().Clone()
		hub.Recover(r)
		hub.Flush(flushTime)
		for _, opt := range opts {
			opt.OnPanic(r)
		}
	}
}

func Recover() {
	if r := recover(); r != nil {
		hub := sentry.CurrentHub().Clone()
		hub.Recover(r)
		hub.Flush(flushTime)
	}
}

func Go(fn func()) {
	go func() {
		defer Recover()
		fn()
	}()
}
