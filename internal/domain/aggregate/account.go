package aggregate

import (
	"github.com/illusory-server/accounts/internal/domain/entity"
	"github.com/illusory-server/accounts/internal/domain/event"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/pkg/errors"
	"time"
)

type Account struct {
	account *entity.Account
	events  []event.Event
}

const (
	DefaultEventCapacity = 4
)

func NewAccount(account *entity.Account) (*Account, error) {
	if err := account.Validate(); err != nil {
		return nil, errors.Wrap(err, "[Account] account.Validate")
	}
	events := make([]event.Event, 0, DefaultEventCapacity)

	return &Account{
		account: account,
		events:  events,
	}, nil
}

func (a *Account) Account() *entity.Account {
	return a.account
}

func (a *Account) Events() []event.Event {
	return a.events
}

func (a *Account) HasEvents() bool {
	return len(a.events) > 0
}

func (a *Account) AddEvent(event event.Event) {
	a.events = append(a.events, event)
}

func (a *Account) ClearEvent() {
	a.events = a.events[:0]
}

func (a *Account) ChangeNickname(nickname string) error {
	err := a.account.SetNickname(nickname)
	if err != nil {
		return errors.Wrap(err, "[Account] account.SetNickname")
	}

	a.AddEvent(event.AccountChangeNickname{
		ID:       a.account.ID(),
		Nickname: nickname,
		Time:     time.Now(),
	})

	return nil
}

func (a *Account) ChangePassword(password vo.Password) error {
	err := a.account.SetPassword(password)
	if err != nil {
		return errors.Wrap(err, "[Account] account.SetPassword")
	}

	a.AddEvent(event.AccountChangePassword{
		ID:       a.account.ID(),
		Password: password,
		Time:     time.Now(),
	})

	return nil
}
