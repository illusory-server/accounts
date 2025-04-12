package psql

import (
	mockPsql "github.com/illusory-server/accounts/internal/mock/psql"
	"github.com/illusory-server/accounts/pkg/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountQuery_Constructor(t *testing.T) {
	query := &mockPsql.MockQueryExecutor{}
	log := logger.NoopLogger{}

	command, err := NewAccountQuery(log, query)
	assert.NoError(t, err)
	assert.NotZero(t, command)

	command, err = NewAccountQuery(nil, query)
	assert.Error(t, err)

	command, err = NewAccountQuery(nil, nil)
	assert.Error(t, err)
}
