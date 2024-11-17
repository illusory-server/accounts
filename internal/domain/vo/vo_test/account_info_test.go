package vo_test

import (
	"github.com/illusory-server/accounts/internal/domain/vo"
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
		assert.Equal(t, vo.AccountInfo{}, info)

		firstName = "eer0"
		lastName = "k"
		info, err = vo.NewAccountInfo(
			firstName, lastName, email,
		)
		assert.Error(t, err)
		assert.Equal(t, vo.AccountInfo{}, info)

		lastName = "kirov"
		email = "not-email"
		info, err = vo.NewAccountInfo(
			firstName, lastName, email,
		)
		assert.Error(t, err)
		assert.Equal(t, vo.AccountInfo{}, info)
	})
}
