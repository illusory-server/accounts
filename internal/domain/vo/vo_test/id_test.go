package vo_test

import (
	"github.com/google/uuid"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVoId(t *testing.T) {
	t.Run("Should correct constructor id", func(t *testing.T) {
		id := uuid.New().String()
		voID, err := vo.NewID(id)
		assert.NoError(t, err)
		assert.Equal(t, id, voID.Value())
	})

	t.Run("Should error with incorrect id", func(t *testing.T) {
		id := "iccorrect-uuidv4-value"
		voID, err := vo.NewID(id)
		assert.Equal(t, voID, vo.ID{})
		assert.Error(t, err)
	})
}
