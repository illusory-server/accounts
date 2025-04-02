package account

import (
	"time"

	"github.com/illusory-server/accounts/internal/domain/vo"
)

type WithoutPassword struct {
	id        vo.ID
	firstName string
	lastName  string
	email     string
	nickname  string
	role      string
	avatarURL string
	updatedAt time.Time
	createdAt time.Time
}

func (w *WithoutPassword) ID() vo.ID {
	return w.id
}

func (w *WithoutPassword) FirstName() string {
	return w.firstName
}

func (w *WithoutPassword) LastName() string {
	return w.lastName
}

func (w *WithoutPassword) Email() string {
	return w.email
}

func (w *WithoutPassword) Nickname() string {
	return w.nickname
}

func (w *WithoutPassword) Role() string {
	return w.role
}

func (w *WithoutPassword) AvatarURL() string {
	return w.avatarURL
}

func (w *WithoutPassword) UpdatedAt() time.Time {
	return w.updatedAt
}

func (w *WithoutPassword) CreatedAt() time.Time {
	return w.createdAt
}
