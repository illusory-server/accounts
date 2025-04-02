package event

import (
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccountChangeNickname(t *testing.T) {
	// Подготовка тестовых данных
	id, _ := vo.NewID("550e8400-e29b-41d4-a716-446655440000")
	nickname := "test_nickname"
	timestamp := time.Now()

	// Создание события
	event := NewAccountChangeNickname(id, nickname, timestamp)

	// Проверка полей
	assert.Equal(t, id, event.ID())
	assert.Equal(t, nickname, event.Nickname())
	assert.Equal(t, timestamp, event.Timestamp())
	assert.Equal(t, AccountChangeNicknameType, event.Type())
}

func TestAccountChangeInfo(t *testing.T) {
	// Подготовка тестовых данных
	id, _ := vo.NewID("550e8400-e29b-41d4-a716-446655440000")
	info, _ := vo.NewAccountInfo("Иван", "Иванов", "ivan@example.com")
	timestamp := time.Now()

	// Создание события
	event := NewAccountChangeInfo(id, info, timestamp)

	// Проверка полей
	assert.Equal(t, id, event.ID())
	assert.Equal(t, info, event.Info())
	assert.Equal(t, timestamp, event.Timestamp())
	assert.Equal(t, AccountChangeInfoType, event.Type())
}

func TestAccountChangePassword(t *testing.T) {
	// Подготовка тестовых данных
	id, _ := vo.NewID("550e8400-e29b-41d4-a716-446655440000")
	password, _ := vo.NewPassword("secure_password123")
	timestamp := time.Now()

	// Создание события
	event := NewAccountChangePassword(id, password, timestamp)

	// Проверка полей
	assert.Equal(t, id, event.ID())
	assert.Equal(t, password, event.Password())
	assert.Equal(t, timestamp, event.Timestamp())
	assert.Equal(t, AccountChangePasswordType, event.Type())
}

func TestAccountChangeEmail(t *testing.T) {
	// Подготовка тестовых данных
	id, _ := vo.NewID("550e8400-e29b-41d4-a716-446655440000")
	email := "new@example.com"
	timestamp := time.Now()

	// Создание события
	event := NewAccountChangeEmail(id, email, timestamp)

	// Проверка полей
	assert.Equal(t, id, event.ID())
	assert.Equal(t, email, event.Email())
	assert.Equal(t, timestamp, event.Timestamp())
	assert.Equal(t, AccountChangeEmailType, event.Type())
}

func TestAccountChangeRole(t *testing.T) {
	// Подготовка тестовых данных
	id, _ := vo.NewID("550e8400-e29b-41d4-a716-446655440000")
	role, _ := vo.NewRole(vo.RoleAdmin)
	timestamp := time.Now()

	// Создание события
	event := NewAccountChangeRole(id, role, timestamp)

	// Проверка полей
	assert.Equal(t, id, event.ID())
	assert.Equal(t, role, event.Role())
	assert.Equal(t, timestamp, event.Timestamp())
	assert.Equal(t, AccountChangeRoleType, event.Type())
}

func TestAccountChangeAvatarLink(t *testing.T) {
	// Подготовка тестовых данных
	id, _ := vo.NewID("550e8400-e29b-41d4-a716-446655440000")
	link, _ := vo.NewLink("https://example.com/avatar.jpg")
	timestamp := time.Now()

	// Создание события
	event := NewAccountChangeAvatarLink(id, link, timestamp)

	// Проверка полей
	assert.Equal(t, id, event.ID())
	assert.Equal(t, link, event.AvatarLink())
	assert.Equal(t, timestamp, event.Timestamp())
	assert.Equal(t, AccountChangeAvatarLinkType, event.Type())
}
