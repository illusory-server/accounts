package ayaka

import (
	"context"
	"fmt"
	ayaka "github.com/illusory-server/accounts/pkg/core"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
	"testing"
	"time"
)

const (
	initEnd        = "init end"
	initEndWithCtx = "init end with ctx done"
	runEnd         = "run end"
	runEndWithCtx  = "run end with ctx done"
)

type correctJob struct {
	initDuration time.Duration
	runDuration  time.Duration
	errInit      error
	errRun       error
	panicInit    string
	panicRun     string
}

func (c correctJob) Init(ctx context.Context, container *dig.Container) error {
	var (
		logger ayaka.Logger
	)
	err := container.Invoke(func(loggerDI ayaka.Logger) {
		logger = loggerDI
	})
	if err != nil {
		return err
	}

	if c.panicInit != "" {
		panic(c.panicInit)
	}
	if c.errInit != nil {
		return c.errInit
	}

	t := time.NewTimer(c.initDuration)
	select {
	case <-t.C:
		logger.Debug(ctx, initEnd, nil)
		return nil
	case <-ctx.Done():
		logger.Debug(ctx, initEndWithCtx, nil)
		return ctx.Err()
	}
}

func (c correctJob) Run(ctx context.Context, container *dig.Container) error {
	var (
		logger ayaka.Logger
	)
	err := container.Invoke(func(loggerDI ayaka.Logger) {
		logger = loggerDI
	})
	if err != nil {
		return err
	}

	if c.panicRun != "" {
		panic(c.panicRun)
	}
	if c.errRun != nil {
		return c.errRun
	}

	t := time.NewTimer(c.runDuration)
	select {
	case <-t.C:
		logger.Debug(ctx, runEnd, nil)
		return nil
	case <-ctx.Done():
		logger.Debug(ctx, runEndWithCtx, nil)
		return ctx.Err()
	}
}

func TestSingleJob(t *testing.T) {
	t.Run("Should correct init and run job", func(t *testing.T) {
		t.Parallel()
		logger := newTestLogger()

		app := ayaka.NewApp(&ayaka.Options{
			Name:        "my-app",
			Description: "my-app description testing",
			Version:     "1.0.0",
			Logger:      logger,
		}).WithConfig(&ayaka.Config{
			StartTimeout:    time.Second * 5,
			GracefulTimeout: time.Second * 5,
		}).WithJob(ayaka.JobEntry{
			Key: "my-test-job",
			Job: &correctJob{
				initDuration: time.Second * 1,
				runDuration:  time.Second * 1,
			},
		})

		err := app.Start()
		assert.NoError(t, err)
		assert.NoError(t, app.Err())

		// TODO asserting logger infos
	})

	t.Run("Should correct error handle init job", func(t *testing.T) {
		t.Parallel()
		logger := newTestLogger()
		myErr := errors.New("my error")

		app := ayaka.NewApp(&ayaka.Options{
			Name:        "my-app",
			Description: "my-app description testing",
			Version:     "1.0.0",
			Logger:      logger,
		}).WithConfig(&ayaka.Config{
			StartTimeout:    time.Second * 5,
			GracefulTimeout: time.Second * 5,
		}).WithJob(ayaka.JobEntry{
			Key: "my-test-job-1",
			Job: &correctJob{
				errInit: myErr,
			},
		})

		err := app.Start()
		assert.Error(t, err)
		assert.Equal(t,
			fmt.Errorf(ayaka.FormatErrJobInitFailed, "my-test-job-1", myErr).Error(),
			errors.Cause(err).Error(),
		)

		logger.messages = logger.messages[2:]
		logger.levels = logger.levels[2:]
		logger.infos = logger.infos[2:]

		assert.Equal(t, []string{ayaka.LogMessageInitError}, logger.messages)
		assert.Equal(t, []string{"error"}, logger.levels)
		assert.Equal(t, []map[string]any{
			{
				ayaka.LogKeyInfoError: myErr.Error(),
				ayaka.LogKeyInfoKey:   "my-test-job-1",
			},
		}, logger.infos)
	})

	t.Run("Should correct panic handle init job", func(t *testing.T) {
		t.Parallel()
		logger := newTestLogger()
		panicMessage := "panic init haha!!!"

		app := ayaka.NewApp(&ayaka.Options{
			Name:        "my-app",
			Description: "my-app description testing",
			Version:     "1.0.0",
			Logger:      logger,
		}).WithConfig(&ayaka.Config{
			StartTimeout:    time.Second * 5,
			GracefulTimeout: time.Second * 5,
		}).WithJob(ayaka.JobEntry{
			Key: "my-test-job-1",
			Job: &correctJob{
				panicInit: panicMessage,
			},
		})

		err := app.Start()
		assert.Error(t, err)
		assert.Equal(t,
			fmt.Errorf(ayaka.FormatErrJobInitPanic, "my-test-job-1", panicMessage).Error(),
			errors.Cause(err).Error(),
		)

		logger.messages = logger.messages[2:]
		logger.levels = logger.levels[2:]
		logger.infos = logger.infos[2:]

		assert.Equal(t, []string{ayaka.LogMessageInitPanic}, logger.messages)
		assert.Equal(t, []string{"error"}, logger.levels)
		assert.Equal(t, []map[string]any{
			{
				ayaka.LogKeyInfoPanic: panicMessage,
				ayaka.LogKeyInfoKey:   "my-test-job-1",
			},
		}, logger.infos)
	})

	t.Run("Should correct error handle run job", func(t *testing.T) {
		t.Parallel()
		logger := newTestLogger()
		myErr := errors.New("my error")

		app := ayaka.NewApp(&ayaka.Options{
			Name:        "my-app",
			Description: "my-app description testing",
			Version:     "1.0.0",
			Logger:      logger,
		}).WithConfig(&ayaka.Config{
			StartTimeout:    time.Second * 5,
			GracefulTimeout: time.Second * 5,
		}).WithJob(ayaka.JobEntry{
			Key: "my-test-job-1",
			Job: &correctJob{
				errRun: myErr,
			},
		})

		err := app.Start()
		assert.Error(t, err)
		assert.Equal(t,
			fmt.Errorf(ayaka.FormatErrJobRunFailed, "my-test-job-1", myErr).Error(),
			errors.Cause(err).Error(),
		)

		logger.messages = logger.messages[4:]
		logger.levels = logger.levels[4:]
		logger.infos = logger.infos[4:]

		assert.Equal(t, []string{ayaka.LogMessageRunError}, logger.messages)
		assert.Equal(t, []string{"error"}, logger.levels)
		assert.Equal(t, []map[string]any{
			{
				ayaka.LogKeyInfoError: myErr.Error(),
				ayaka.LogKeyInfoKey:   "my-test-job-1",
			},
		}, logger.infos)
	})

	t.Run("Should correct panic handle run job", func(t *testing.T) {
		t.Parallel()
		logger := newTestLogger()
		panicMessage := "panic run haha!!!"

		app := ayaka.NewApp(&ayaka.Options{
			Name:        "my-app",
			Description: "my-app description testing",
			Version:     "1.0.0",
			Logger:      logger,
		}).WithConfig(&ayaka.Config{
			StartTimeout:    time.Second * 5,
			GracefulTimeout: time.Second * 5,
		}).WithJob(ayaka.JobEntry{
			Key: "my-test-job-1",
			Job: &correctJob{
				panicRun: panicMessage,
			},
		})

		err := app.Start()
		assert.Error(t, err)
		assert.Equal(t,
			fmt.Errorf(ayaka.FormatErrJobRunPanic, "my-test-job-1", panicMessage).Error(),
			errors.Cause(err).Error(),
		)

		logger.messages = logger.messages[4:]
		logger.levels = logger.levels[4:]
		logger.infos = logger.infos[4:]

		assert.Equal(t, []string{ayaka.LogMessageRunPanic}, logger.messages)
		assert.Equal(t, []string{"error"}, logger.levels)
		assert.Equal(t, []map[string]any{
			{
				ayaka.LogKeyInfoPanic: panicMessage,
				ayaka.LogKeyInfoKey:   "my-test-job-1",
			},
		}, logger.infos)
	})
}

func TestMultipleJobs(t *testing.T) {
	t.Run("Should correct init and run jobs", func(t *testing.T) {

	})

	t.Run("Should correct error handle init jobs", func(t *testing.T) {

	})

	t.Run("Should correct error handle run jobs", func(t *testing.T) {

	})

	t.Run("Should correct error handle run jobs", func(t *testing.T) {

	})

	t.Run("Should correct panic handle run jobs", func(t *testing.T) {

	})
}

func TestJobsTimout(t *testing.T) {
	t.Run("Should correct stop init with start timout", func(t *testing.T) {

	})

	t.Run("Should correct graceful timeout init job", func(t *testing.T) {

	})

	t.Run("Should correct graceful timeout run job", func(t *testing.T) {

	})
}
