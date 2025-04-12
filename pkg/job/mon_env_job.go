package job

import (
	"github.com/OddEer0/ayaka/ecosystem"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"time"
)

type (
	MonJobParams struct {
		ReadTimeout    time.Duration
		WriteTimeout   time.Duration
		MaxHeaderBytes int
		Address        string
	}

	MonJobEnvKeys struct {
		Address, Timeout string
	}
)

var (
	DefaultMonJobKeys = &MonJobEnvKeys{
		Address: "MON_ADDRESS",
		Timeout: "MON_TIMEOUT",
	}
)

func NewMonJob[T any](params MonJobParams) (*ecosystem.MonitoringJob[T], error) {
	return ecosystem.NewMonitoringJobBuilder[T]().
		Address(params.Address).
		Build()
}

func getMonJobParamsEnv(keys *MonJobEnvKeys) (MonJobParams, error) {
	addr := os.Getenv(keys.Address)
	if addr == "" {
		return MonJobParams{}, errors.Errorf("environment variable %s is not set", keys.Address)
	}
	timeout := os.Getenv(keys.Timeout)
	if timeout == "" {
		return MonJobParams{}, errors.Errorf("environment variable %s is not set", keys.Timeout)
	}
	requestTimeout, err := strconv.Atoi(timeout)
	if err != nil {
		return MonJobParams{}, errors.Errorf("environment variable %s is not a number", keys.Timeout)
	}

	return MonJobParams{
		Address:        addr,
		ReadTimeout:    time.Duration(requestTimeout) * time.Second,
		WriteTimeout:   time.Duration(requestTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20, //nolint:mnd
	}, nil
}

func NewMonJobEnv[T any](keys *MonJobEnvKeys) (*ecosystem.MonitoringJob[T], error) {
	if keys == nil {
		keys = DefaultMonJobKeys
	}
	params, err := getMonJobParamsEnv(keys)
	if err != nil {
		return nil, errors.Wrap(err, "getMonJobParamsEnv")
	}
	return NewMonJob[T](MonJobParams{
		ReadTimeout:    params.ReadTimeout,
		WriteTimeout:   params.WriteTimeout,
		MaxHeaderBytes: params.MaxHeaderBytes,
		Address:        params.Address,
	})
}

func MustMonJobEnv[T any](keys *MonJobEnvKeys) *ecosystem.MonitoringJob[T] {
	job, err := NewMonJobEnv[T](keys)
	if err != nil {
		panic(err)
	}
	return job
}
