package psql

import (
	mockPsql "github.com/illusory-server/accounts/internal/mock/psql"
	"github.com/illusory-server/accounts/pkg/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountCommand_Constructor(t *testing.T) {
	query := &mockPsql.MockQueryExecutor{}
	log := logger.NoopLogger{}

	command, err := NewAccountCommand(log, query)
	assert.NoError(t, err)
	assert.NotZero(t, command)

	command, err = NewAccountCommand(nil, query)
	assert.Error(t, err)

	command, err = NewAccountCommand(nil, nil)
	assert.Error(t, err)
}
