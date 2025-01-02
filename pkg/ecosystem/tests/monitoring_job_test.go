package grpc_job

import (
	ayaka "github.com/illusory-server/accounts/pkg/core"
	"github.com/illusory-server/accounts/pkg/ecosystem"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestMonitoringJobBuilder(t *testing.T) {
	t.Run("Should correctly build monitoring job", func(t *testing.T) {
		address := "localhost:1000"
		job, err := ecosystem.NewMonitoringJobBuilder().
			Address(address).
			Build()

		assert.NoError(t, err)
		assert.Equal(t, address, job.Address())
	})

	t.Run("Should correctly failed building monitoring job", func(t *testing.T) {
		job, err := ecosystem.NewMonitoringJobBuilder().
			Build()

		assert.Error(t, err)
		assert.Nil(t, job)
	})

	t.Run("Should correctly custom mux", func(t *testing.T) {
		address := "localhost:1000"
		job, err := ecosystem.NewMonitoringJobBuilder().
			Address(address).
			Mux(http.NewServeMux()).
			Build()

		assert.NoError(t, err)
		assert.Equal(t, address, job.Address())
	})
}

func TestMonitoringJobSignature(t *testing.T) {
	address := "localhost:10101"

	job, err := ecosystem.NewMonitoringJobBuilder().
		Address(address).
		Build()

	assert.NoError(t, err)

	ayaka.NewApp(&ayaka.Options{
		Name:        "aya",
		Description: "kekw",
		Version:     "0.0.1",
		Container:   ayaka.NewContainer(ayaka.NoopLogger{}),
	}).WithJob(ayaka.JobEntry{
		Key: "xd",
		Job: job,
	})
}
