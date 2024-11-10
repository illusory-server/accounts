package model

import (
	"github.com/illusory-server/accounts/internal/domain/vo"
	"time"
)

type Account struct {
	Id        vo.ID
	Nickname  string
	FirstName string
	LastName  string
	Email     string
	Role      string
	Password  string
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (a *Account) FullName() string {
	return a.FirstName + a.LastName
}
