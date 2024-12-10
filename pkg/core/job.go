package ayaka

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	LogKeyInfoKey                    = "job_key"
	LogKeyInfoError                  = "job_error"
	LogKeyInfoPanic                  = "job_panic"
	LogMessageInitError              = "job init failed"
	LogMessageInitPanic              = "job init panic"
	LogMessageRunError               = "job run failed"
	LogMessageRunPanic               = "job run panic"
	LogMessageGracefulShotdownFailed = "graceful shotdown failed"
	FormatErrJobInitFailed           = "failed to initialize job '%s': %w"
	FormatErrJobInitPanic            = "panic in initialized job '%s': %v"
	FormatErrJobRunFailed            = "failed run job '%s': %w"
	FormatErrJobRunPanic             = "panic in runned job '%s': %v"
)

func (a *App) initJob(ctx context.Context) error {
	var wg sync.WaitGroup
	ctxStop, cancel := context.WithCancel(ctx)
	defer cancel()

	sErr := newSingleError(func() {
		cancel()
	})
	stopChan := make(chan struct{})

	for key, job := range a.jobs {
		wg.Add(1)
		go func(ctx context.Context, key string, job Job) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					a.logger.Error(ctx, LogMessageInitPanic, map[string]any{
						LogKeyInfoKey:   key,
						LogKeyInfoPanic: r,
					})
					sErr.add(fmt.Errorf(FormatErrJobInitPanic, key, r))
				}
			}()

			if err := job.Init(ctxStop, a.Dependency()); err != nil {
				a.logger.Error(ctx, LogMessageInitError, map[string]any{
					LogKeyInfoKey:   key,
					LogKeyInfoError: err.Error(),
				})
				sErr.add(fmt.Errorf(FormatErrJobInitFailed, key, err))
			}
		}(a.ctx, key, job)
	}

	go func() {
		wg.Wait()
		close(stopChan)
	}()

	select {
	case <-stopChan:
		return sErr.get()
	case <-ctx.Done():
		t := time.NewTimer(a.Config().GracefulTimeout)
		select {
		case <-t.C:
			a.logger.Warn(a.ctx, LogMessageGracefulShotdownFailed, nil)
			return ErrGracefulTimeout
		case <-stopChan:
			return ctx.Err()
		}
	}
}

func (a *App) runJob(ctx context.Context) error {
	var wg sync.WaitGroup
	ctxStop, cancel := context.WithCancel(ctx)
	defer cancel()

	sErr := newSingleError(func() {
		cancel()
	})
	stopChan := make(chan struct{})

	for key, job := range a.jobs {
		wg.Add(1)
		go func(ctx context.Context, key string, job Job) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					a.logger.Error(ctx, LogMessageRunPanic, map[string]any{
						LogKeyInfoKey:   key,
						LogKeyInfoPanic: r,
					})
					sErr.add(fmt.Errorf(FormatErrJobRunPanic, key, r))
				}
			}()

			if err := job.Run(ctxStop, a.Dependency()); err != nil {
				a.logger.Error(ctx, LogMessageRunError, map[string]any{
					LogKeyInfoKey:   key,
					LogKeyInfoError: err.Error(),
				})
				sErr.add(fmt.Errorf(FormatErrJobRunFailed, key, err))
			}
		}(a.ctx, key, job)
	}

	go func() {
		wg.Wait()
		close(stopChan)
	}()

	select {
	case <-stopChan:
		return sErr.get()
	case <-ctx.Done():
		t := time.NewTimer(a.Config().GracefulTimeout)
		select {
		case <-t.C:
			a.logger.Warn(a.ctx, LogMessageGracefulShotdownFailed, nil)
			return ErrGracefulTimeout
		case <-stopChan:
			return ctx.Err()
		}
	}
}
