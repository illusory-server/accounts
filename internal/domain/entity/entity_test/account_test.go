package entity_test

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/illusory-server/accounts/internal/domain"
	"github.com/illusory-server/accounts/internal/domain/entity"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/illusory-server/accounts/pkg/errors/codex"
	"github.com/illusory-server/accounts/pkg/errors/errx"
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
		assert.Equal(t, codex.InvalidArgument, errx.Code(err))
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
		assert.Equal(t, codex.InvalidArgument, errx.Code(err))

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
		assert.Equal(t, codex.InvalidArgument, errx.Code(err))

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
		assert.Equal(t, codex.InvalidArgument, errx.Code(err))

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
		assert.Equal(t, codex.InvalidArgument, errx.Code(err))

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
		assert.Equal(t, codex.InvalidArgument, errx.Code(err))

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
		assert.Equal(t, codex.InvalidArgument, errx.Code(err))
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
		assert.Equal(t, codex.InvalidArgument, errx.Code(err))
		assert.Equal(t, newPass, acc.Password())

		newInfo, err := vo.NewAccountInfo("eer0-2", "kirov-2", "kuru2@gmail.com")
		assert.NoError(t, err)
		err = acc.SetInfo(newInfo)
		assert.NoError(t, err)
		assert.Equal(t, newInfo, acc.Info())
		err = acc.SetInfo(vo.AccountInfo{})
		assert.Error(t, err)
		assert.Equal(t, codex.InvalidArgument, errx.Code(err))
		assert.Equal(t, newInfo, acc.Info())

		newNick := "new-eer0"
		err = acc.SetNickname(newNick)
		assert.NoError(t, err)
		assert.Equal(t, newNick, acc.Nickname())
		err = acc.SetNickname("s")
		assert.Error(t, err)
		assert.Equal(t, codex.InvalidArgument, errx.Code(err))
		assert.Equal(t, newNick, acc.Nickname())

		newUpdatedTime := time.Now()
		err = acc.SetUpdatedAt(newUpdatedTime)
		assert.NoError(t, err)
		assert.Equal(t, newUpdatedTime, acc.UpdatedAt())
		err = acc.SetUpdatedAt(time.Now().Add(1 * time.Hour))
		assert.Error(t, err)
		assert.Equal(t, codex.InvalidArgument, errx.Code(err))
		assert.Equal(t, newUpdatedTime, acc.UpdatedAt())

		avatar, err := vo.NewLink("https://joska.com/5432435")
		assert.NoError(t, err)
		err = acc.SetAvatarLink(avatar)
		assert.NoError(t, err)
		assert.Equal(t, avatar, acc.AvatarLink().ValueOrDefault(vo.Link{}))

		err = acc.SetAvatarLink(vo.Link{})
		assert.Error(t, err)
		assert.Equal(t, codex.InvalidArgument, errx.Code(err))

		roleSet, err := vo.NewRole(vo.RoleUser)
		assert.NoError(t, err)
		assert.Equal(t, vo.RoleSuperAdmin, acc.Role().Value())
		err = acc.SetRole(roleSet)
		assert.NoError(t, err)
		assert.Equal(t, roleSet.Value(), acc.Role().Value())

		err = acc.SetRole(vo.Role{})
		assert.Error(t, err)
		assert.Equal(t, codex.InvalidArgument, errx.Code(err))
		assert.Equal(t, roleSet.Value(), acc.Role().Value())
	})

	t.Run("Should correct marshal", func(t *testing.T) {
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

		s, err := json.Marshal(acc)
		assert.NoError(t, err)
		var m map[string]interface{}
		err = json.Unmarshal(s, &m)
		assert.NoError(t, err)

		expected := map[string]interface{}{
			"id": id.Value(),
			"info": map[string]interface{}{
				"first_name": info.FirstName(),
				"last_name":  info.LastName(),
				"email":      info.Email(),
			},
			"role":        role.Value(),
			"nickname":    nick,
			"updated_at":  updatedTime,
			"avatar_link": "",
			"created_at":  createdTime,
		}

		expectedBytes, err := json.Marshal(expected)
		assert.NoError(t, err)
		var expectedMap map[string]interface{}
		err = json.Unmarshal(expectedBytes, &expectedMap)
		assert.NoError(t, err)
		assert.Equal(t, expectedMap, m)
	})

	t.Run("Should correct getter", func(t *testing.T) {
		var acc *entity.Account
		assert.Equal(t, vo.ID{}, acc.ID())
		assert.Equal(t, vo.AccountInfo{}, acc.Info())
		assert.Equal(t, vo.Password{}, acc.Password())
		assert.Equal(t, vo.Role{}, acc.Role())
		assert.Equal(t, "", acc.Nickname())
		assert.Equal(t, domain.NewEmptyOptional[vo.Link](), acc.AvatarLink())
		assert.Equal(t, time.Time{}, acc.UpdatedAt())
		assert.Equal(t, time.Time{}, acc.CreatedAt())

	})
}
