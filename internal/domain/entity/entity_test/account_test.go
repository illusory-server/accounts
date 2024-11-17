package entity_test

import (
	"github.com/google/uuid"
	"github.com/illusory-server/accounts/internal/domain/entity"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEntityAccount(t *testing.T) {
	id, err := vo.NewID(uuid.New().String())
	assert.NoError(t, err)
	info, err := vo.NewAccountInfo("eer0", "kirov", "kuru@gmail.com")
	assert.NoError(t, err)
	role, err := vo.NewRole(vo.RoleSuperAdmin)
	assert.NoError(t, err)
	pass, err := vo.NewPassword("correct-pass123")
	assert.NoError(t, err)
	nick := "eer0"
	updatedTime := time.Now().Add(-1 * time.Hour)
	createdTime := time.Now().Add(-2 * time.Hour)
	t.Run("Should correct constructor create account", func(t *testing.T) {
		acc, err := entity.NewAccount(
			id,
			info,
			role,
			nick,
			pass,
			updatedTime,
			createdTime,
		)
		assert.NoError(t, err)
		assert.Equal(t, id, acc.ID())
		assert.Equal(t, info, acc.Info())
		assert.Equal(t, role, acc.Role())
		assert.Equal(t, nick, acc.Nickname())
		assert.Equal(t, pass, acc.Password())
		assert.Equal(t, updatedTime, acc.UpdatedAt())
		assert.Equal(t, createdTime, acc.CreatedAt())
	})

	t.Run("Should incorrect constructor create account", func(t *testing.T) {
		acc, err := entity.NewAccount(
			vo.ID{},
			info,
			role,
			nick,
			pass,
			updatedTime,
			createdTime,
		)
		assert.Nil(t, acc)
		assert.Error(t, err)

		acc, err = entity.NewAccount(
			id,
			vo.AccountInfo{},
			role,
			nick,
			pass,
			updatedTime,
			createdTime,
		)
		assert.Nil(t, acc)
		assert.Error(t, err)

		acc, err = entity.NewAccount(
			id,
			info,
			vo.Role{},
			nick,
			pass,
			updatedTime,
			createdTime,
		)
		assert.Nil(t, acc)
		assert.Error(t, err)

		acc, err = entity.NewAccount(
			id,
			info,
			role,
			nick,
			vo.Password{},
			updatedTime,
			createdTime,
		)
		assert.Nil(t, acc)
		assert.Error(t, err)

		acc, err = entity.NewAccount(
			id,
			info,
			role,
			"e",
			pass,
			updatedTime,
			createdTime,
		)
		assert.Nil(t, acc)
		assert.Error(t, err)

		incorrectTime := time.Now().Add(1 * time.Hour)
		acc, err = entity.NewAccount(
			id,
			info,
			role,
			nick,
			pass,
			incorrectTime,
			createdTime,
		)
		assert.Nil(t, acc)
		assert.Error(t, err)

		acc, err = entity.NewAccount(
			id,
			info,
			role,
			nick,
			pass,
			updatedTime,
			incorrectTime,
		)
		assert.Nil(t, acc)
		assert.Error(t, err)
	})

	t.Run("Should correct setter", func(t *testing.T) {
		acc, err := entity.NewAccount(
			id,
			info,
			role,
			nick,
			pass,
			updatedTime,
			createdTime,
		)
		assert.NoError(t, err)

		newPass, err := vo.NewPassword("correct_pass_xd2")
		assert.NoError(t, err)
		err = acc.SetPassword(newPass)
		assert.NoError(t, err)
		assert.Equal(t, newPass, acc.Password())
		err = acc.SetPassword(vo.Password{})
		assert.Error(t, err)
		assert.Equal(t, newPass, acc.Password())

		newInfo, err := vo.NewAccountInfo("eer0-2", "kirov-2", "kuru2@gmail.com")
		assert.NoError(t, err)
		err = acc.SetInfo(newInfo)
		assert.NoError(t, err)
		assert.Equal(t, newInfo, acc.Info())
		err = acc.SetInfo(vo.AccountInfo{})
		assert.Error(t, err)
		assert.Equal(t, newInfo, acc.Info())

		newNick := "new-eer0"
		err = acc.SetNickname(newNick)
		assert.NoError(t, err)
		assert.Equal(t, newNick, acc.Nickname())
		err = acc.SetNickname("s")
		assert.Error(t, err)
		assert.Equal(t, newNick, acc.Nickname())

		newUpdatedTime := time.Now()
		err = acc.SetUpdatedAt(newUpdatedTime)
		assert.NoError(t, err)
		assert.Equal(t, newUpdatedTime, acc.UpdatedAt())
		err = acc.SetUpdatedAt(time.Now().Add(1 * time.Hour))
		assert.Error(t, err)
		assert.Equal(t, newUpdatedTime, acc.UpdatedAt())
	})
}