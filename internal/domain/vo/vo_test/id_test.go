package vo_test

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/illusory-server/accounts/pkg/errors/codex"
	"github.com/illusory-server/accounts/pkg/errors/errx"
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
		assert.Equal(t, codex.InvalidArgument, errx.Code(err))
		assert.Error(t, err)
	})

	t.Run("Should marshal", func(t *testing.T) {
		id := uuid.New().String()
		voID, err := vo.NewID(id)
		assert.NoError(t, err)
		s, err := json.Marshal(voID)
		assert.NoError(t, err)
		assert.Equal(t, "\""+id+"\"", string(s))
	})
}
