package grpc_job

import (
	"context"
	"fmt"
	ayaka "github.com/illusory-server/accounts/pkg/core"
	"github.com/illusory-server/accounts/pkg/ecosystem"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

type printer struct {
	PrintString string
}

func (p *printer) Printf(format string, args ...interface{}) {
	p.PrintString = fmt.Sprintf(format, args...)
}

type activateJob struct {
	initCount int
	runCount  int
}

func (a *activateJob) Init(ctx context.Context, container ayaka.Container) error {
	a.initCount++
	return nil
}

func (a *activateJob) Run(ctx context.Context, container ayaka.Container) error {
	a.runCount++
	return nil
}

func TestStartWithCli(t *testing.T) {
	orig := os.Args
	defer func() { os.Args = orig }()
	job := &activateJob{}

	appVersion := "1.0.0"
	app := ayaka.NewApp(&ayaka.Options{
		Name:        "TestStartWithCli",
		Description: "TestStartWithCli case",
		Version:     appVersion,
		Container:   ayaka.NewContainer(ayaka.NoopLogger{}),
	}).
		WithConfig(&ayaka.Config{
			StartTimeout:    500 * time.Millisecond,
			GracefulTimeout: 5 * time.Second,
		}).
		WithJob(ayaka.JobEntry{
			Key: "test",
			Job: job,
		})
	p := &printer{}

	os.Args = []string{"program_name", "version"}
	err := ecosystem.StartWithCli(app, p)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("VERSION: %s\n", appVersion), p.PrintString)

	os.Args = []string{"program_name", "help"}
	err = ecosystem.StartWithCli(app, p)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%s\n", ecosystem.CliHelpString), p.PrintString)

	os.Args = []string{"program_name", "run"}
	err = ecosystem.StartWithCli(app, nil)
	assert.NoError(t, err)
	assert.Equal(t, 1, job.initCount)
	assert.Equal(t, 1, job.runCount)
}
