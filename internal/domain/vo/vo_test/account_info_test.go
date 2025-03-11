package vo_test

import (
	"encoding/json"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/illusory-server/accounts/pkg/errors/codes"
	"github.com/illusory-server/accounts/pkg/errors/errx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVoAccountInfo(t *testing.T) {
	t.Run("Should correct constructor account info", func(t *testing.T) {
		firstName := "eer0"
		lastName := "kirov"
		email := "kuru@gmail.com"
		info, err := vo.NewAccountInfo(
			firstName, lastName, email,
		)
		assert.NoError(t, err)
		assert.Equal(t, firstName, info.FirstName())
		assert.Equal(t, lastName, info.LastName())
		assert.Equal(t, email, info.Email())

		assert.Equal(t, firstName+" "+lastName, info.FullName())
	})

	t.Run("Should error with incorrect account info", func(t *testing.T) {
		firstName := "e"
		lastName := "kirov"
		email := "kuru@gmail.com"
		info, err := vo.NewAccountInfo(
			firstName, lastName, email,
		)
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, errx.Code(err))
		assert.Equal(t, vo.AccountInfo{}, info)

		firstName = "eer0"
		lastName = "k"
		info, err = vo.NewAccountInfo(
			firstName, lastName, email,
		)
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, errx.Code(err))
		assert.Equal(t, vo.AccountInfo{}, info)

		lastName = "kirov"
		email = "not-email"
		info, err = vo.NewAccountInfo(
			firstName, lastName, email,
		)
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, errx.Code(err))
		assert.Equal(t, vo.AccountInfo{}, info)
	})

	t.Run("Should marshal json", func(t *testing.T) {
		firstName := "eer0"
		lastName := "kirov"
		email := "kuru@gmail.com"
		info, err := vo.NewAccountInfo(
			firstName, lastName, email,
		)
		assert.NoError(t, err)

		bytes, err := json.Marshal(info)
		assert.NoError(t, err)

		var m map[string]interface{}
		err = json.Unmarshal(bytes, &m)
		assert.NoError(t, err)

		expected := map[string]interface{}{
			"first_name": firstName,
			"last_name":  lastName,
			"email":      email,
		}
		assert.Equal(t, expected, m)
	})
}
