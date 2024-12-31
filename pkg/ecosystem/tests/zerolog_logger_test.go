package grpc_job

import (
	"github.com/illusory-server/accounts/pkg/ecosystem"
	"github.com/rs/zerolog"
	"testing"
)

func TestAppLoggerFromZerolog(t *testing.T) {
	output := &testOut{}
	logger := zerolog.New(output)

	log := ecosystem.NewAppLoggerWithZerolog(&logger)

	testAppLogger(t, log, output, "message")
}
