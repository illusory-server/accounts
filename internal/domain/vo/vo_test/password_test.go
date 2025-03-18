package vo

import (
	"encoding/json"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/illusory-server/accounts/pkg/errors/codes"
	"github.com/illusory-server/accounts/pkg/errors/errx"
	"github.com/stretchr/testify/assert"
	"testing"
)

type in struct {
	Field vo.Password `json:"field"`
}

func TestVoPassword(t *testing.T) {
	t.Run("Should correct constructor", func(t *testing.T) {
		correctPass := "correct123"
		pass, err := vo.NewPassword(correctPass)
		assert.NoError(t, err)
		assert.Equal(t, correctPass, pass.Value())
	})

	t.Run("Should return error for invalid password", func(t *testing.T) {
		pass, err := vo.NewPassword("pass")
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, errx.Code(err))
		assert.Equal(t, vo.Password{}, pass)

		pass, err = vo.NewPassword("wrong")
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, errx.Code(err))
		assert.Equal(t, vo.Password{}, pass)
	})

	t.Run("Should correct marshal json", func(t *testing.T) {
		correctPass := "correct123"
		pass, err := vo.NewPassword(correctPass)
		assert.NoError(t, err)
		js, err := json.Marshal(in{Field: pass})
		assert.Equal(t, "{\"field\":\"secret-value\"}", string(js))
	})
}
