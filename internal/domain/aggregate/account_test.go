package aggregate

import (
	"github.com/illusory-server/accounts/internal/domain/entity"
	"github.com/illusory-server/accounts/internal/domain/event"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/illusory-server/accounts/pkg/errors/codex"
	"github.com/illusory-server/accounts/pkg/errors/errx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func createTestAccount(t *testing.T) *Account {
	// Подготовка тестовых данных
	id, err := vo.NewID("550e8400-e29b-41d4-a716-446655440000")
	require.NoError(t, err)

	info, err := vo.NewAccountInfo("Иван", "Иванов", "test@example.com")
	require.NoError(t, err)

	role, err := vo.NewRole(vo.RoleUser)
	require.NoError(t, err)

	hashed, err := bcrypt.GenerateFromPassword([]byte("secure_password123"), bcrypt.MinCost)
	password, err := vo.NewPassword(string(hashed))
	require.NoError(t, err)

	now := time.Now()
	acc, err := entity.NewAccount(
		id,
		info,
		role,
		"test_nickname",
		password,
		now,
		now,
	)
	require.NoError(t, err)

	aggregate, err := NewAccount(acc)
	require.NoError(t, err)

	return aggregate
}

func TestNewAccount(t *testing.T) {
	acc := createTestAccount(t)
	assert.NotNil(t, acc)
	assert.Equal(t, 0, len(acc.Events()))
	assert.False(t, acc.HasEvents())

	res, err := NewAccount(&entity.Account{})
	assert.Error(t, err)
	assert.Nil(t, res)
	c := errx.Code(err)
	assert.Equal(t, codex.InvalidArgument, c)
}

func TestAccount_ChangeNickname(t *testing.T) {
	acc := createTestAccount(t)
	newNickname := "new_nickname"

	err := acc.ChangeNickname(newNickname)
	require.NoError(t, err)

	assert.Equal(t, newNickname, acc.Account().Nickname())

	assert.True(t, acc.HasEvents())
	assert.Equal(t, 1, len(acc.Events()))
	e := acc.Events()[0]
	assert.Equal(t, event.AccountChangeNicknameType, e.Type())

	err = acc.ChangeNickname("")
	assert.Error(t, err)
	assert.Equal(t, codex.InvalidArgument, errx.Code(err))

	assert.Equal(t, newNickname, acc.Account().Nickname())
	assert.Equal(t, 1, len(acc.Events()))
}

func TestAccount_ComparePassword(t *testing.T) {
	acc := createTestAccount(t)

	err := acc.ComparePassword("secure_password123")
	assert.NoError(t, err)

	err = acc.ComparePassword("wrong_password")
	assert.Error(t, err)

	err = acc.ComparePassword("wrong_password")
	assert.Error(t, err)
	assert.Equal(t, codex.InvalidArgument, errx.Code(err))
}

func TestAccount_ChangeAccountInfo(t *testing.T) {
	acc := createTestAccount(t)

	newInfo, err := vo.NewAccountInfo("Петр", "Петров", "petr@example.com")
	require.NoError(t, err)

	err = acc.ChangeAccountInfo(newInfo)
	require.NoError(t, err)

	assert.Equal(t, newInfo, acc.Account().Info())

	assert.True(t, acc.HasEvents())
	assert.Equal(t, 1, len(acc.Events()))
	e := acc.Events()[0]
	assert.Equal(t, event.AccountChangeInfoType, e.Type())

	err = acc.ChangeAccountInfo(vo.AccountInfo{})
	assert.Error(t, err)
	assert.Equal(t, codex.InvalidArgument, errx.Code(err))

	assert.Equal(t, newInfo, acc.Account().Info())
	assert.Equal(t, 1, len(acc.Events()))
	e = acc.Events()[0]
	assert.Equal(t, event.AccountChangeInfoType, e.Type())
}

func TestAccount_ChangeRole(t *testing.T) {
	acc := createTestAccount(t)

	newRole, err := vo.NewRole(vo.RoleAdmin)
	require.NoError(t, err)

	err = acc.ChangeRole(newRole)
	require.NoError(t, err)

	assert.Equal(t, newRole, acc.Account().Role())

	assert.True(t, acc.HasEvents())
	assert.Equal(t, 1, len(acc.Events()))
	e := acc.Events()[0]
	assert.Equal(t, event.AccountChangeRoleType, e.Type())

	err = acc.ChangeRole(vo.Role{})
	assert.Error(t, err)
	assert.Equal(t, codex.InvalidArgument, errx.Code(err))

	assert.Equal(t, newRole, acc.Account().Role())
	assert.Equal(t, 1, len(acc.Events()))
	e = acc.Events()[0]
	assert.Equal(t, event.AccountChangeRoleType, e.Type())
}

func TestAccount_ChangeAvatarLink(t *testing.T) {
	acc := createTestAccount(t)

	newLink, err := vo.NewLink("https://example.com/new-avatar.jpg")
	require.NoError(t, err)

	err = acc.ChangeAvatarLink(newLink)
	require.NoError(t, err)

	avatarLink, err := acc.Account().AvatarLink().Value()
	require.NoError(t, err)
	assert.Equal(t, newLink, avatarLink)

	assert.True(t, acc.HasEvents())
	assert.Equal(t, 1, len(acc.Events()))
	e := acc.Events()[0]
	assert.Equal(t, event.AccountChangeAvatarLinkType, e.Type())

	err = acc.ChangeAvatarLink(vo.Link{})
	assert.Error(t, err)
	assert.Equal(t, codex.InvalidArgument, errx.Code(err))

	avatarLink, err = acc.Account().AvatarLink().Value()
	require.NoError(t, err)
	assert.Equal(t, newLink, avatarLink)

	assert.Equal(t, 1, len(acc.Events()))
	e = acc.Events()[0]
	assert.Equal(t, event.AccountChangeAvatarLinkType, e.Type())
}

func TestAccount_ChangeEmail(t *testing.T) {
	acc := createTestAccount(t)
	email := "new_email@gmail.com"
	newEmail, err := vo.NewAccountInfo(acc.account.Info().FirstName(), acc.account.Info().LastName(), email)
	assert.NoError(t, err)
	err = acc.ChangeEmail(newEmail)
	assert.NoError(t, err)

	assert.Equal(t, email, acc.account.Info().Email())

	assert.True(t, acc.HasEvents())
	assert.Equal(t, 1, len(acc.Events()))
	e := acc.Events()[0]
	assert.Equal(t, event.AccountChangeEmailType, e.Type())

	err = acc.ChangeEmail(vo.AccountInfo{})
	assert.Error(t, err)
	assert.Equal(t, codex.InvalidArgument, errx.Code(err))

	assert.Equal(t, email, acc.account.Info().Email())

	assert.Equal(t, 1, len(acc.Events()))
	e = acc.Events()[0]
	assert.Equal(t, event.AccountChangeEmailType, e.Type())
}

func TestAccount_ChangePassword(t *testing.T) {
	acc := createTestAccount(t)

	newPassword, err := vo.NewPassword("new_secure_password123")
	require.NoError(t, err)

	err = acc.ChangePassword(newPassword)
	require.NoError(t, err)

	assert.Equal(t, newPassword, acc.Account().Password())

	assert.True(t, acc.HasEvents())
	assert.Equal(t, 1, len(acc.Events()))
	e := acc.Events()[0]
	assert.Equal(t, event.AccountChangePasswordType, e.Type())

	err = acc.ChangePassword(vo.Password{})
	assert.Error(t, err)
	assert.Equal(t, codex.InvalidArgument, errx.Code(err))

	assert.Equal(t, newPassword, acc.Account().Password())
	assert.Equal(t, 1, len(acc.Events()))
	e = acc.Events()[0]
	assert.Equal(t, event.AccountChangePasswordType, e.Type())
}

func TestAccount_ClearEvents(t *testing.T) {
	acc := createTestAccount(t)

	newNickname := "new_nickname"
	err := acc.ChangeNickname(newNickname)
	require.NoError(t, err)

	newRole, err := vo.NewRole(vo.RoleAdmin)
	require.NoError(t, err)
	err = acc.ChangeRole(newRole)
	require.NoError(t, err)

	assert.True(t, acc.HasEvents())
	assert.Equal(t, 2, len(acc.Events()))

	acc.ClearEvent()

	assert.False(t, acc.HasEvents())
	assert.Equal(t, 0, len(acc.Events()))
}
