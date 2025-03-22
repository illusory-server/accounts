package aggregate

import (
	"github.com/illusory-server/accounts/internal/domain/entity"
	"github.com/illusory-server/accounts/internal/domain/event"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/illusory-server/accounts/pkg/errors/codex"
	"github.com/illusory-server/accounts/pkg/errors/errx"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
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
		return nil, errx.WrapWithCode(err, codex.InvalidArgument, "[Account] account.Validate")
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

	a.AddEvent(event.NewAccountChangeNickname(a.account.ID(), nickname, time.Now()))

	return nil
}

func (a *Account) ComparePassword(password string) error {
	current := a.account.Password().Value()
	err := bcrypt.CompareHashAndPassword([]byte(current), []byte(password))
	if err != nil {
		return errx.WrapWithCode(err, codex.InvalidArgument, "[Account] bcrypt.CompareHashAndPassword")
	}
	return nil
}

func (a *Account) ChangeAccountInfo(info vo.AccountInfo) error {
	err := a.account.SetInfo(info)
	if err != nil {
		return errors.Wrap(err, "[Account] account.SetInfo")
	}

	a.AddEvent(event.NewAccountChangeInfo(a.account.ID(), info, time.Now()))

	return nil
}

func (a *Account) ChangeEmail(info vo.AccountInfo) error {
	err := a.account.SetInfo(info)
	if err != nil {
		return errors.Wrap(err, "[Account] account.SetInfo")
	}

	a.AddEvent(event.NewAccountChangeEmail(a.account.ID(), info.Email(), time.Now()))

	return nil
}

func (a *Account) ChangeRole(role vo.Role) error {
	err := a.account.SetRole(role)
	if err != nil {
		return errors.Wrap(err, "[Account] account.SetRole")
	}

	a.AddEvent(event.NewAccountChangeRole(a.account.ID(), role, time.Now()))

	return nil
}

func (a *Account) ChangeAvatarLink(link vo.Link) error {
	err := a.account.SetAvatarLink(link)
	if err != nil {
		return errors.Wrap(err, "[Account] account.SetAvatarLink")
	}

	a.AddEvent(event.NewAccountChangeAvatarLink(a.account.ID(), link, time.Now()))

	return nil
}

func (a *Account) ChangePassword(password vo.Password) error {
	err := a.account.SetPassword(password)
	if err != nil {
		return errors.Wrap(err, "[Account] account.SetPassword")
	}

	a.AddEvent(event.NewAccountChangePassword(a.account.ID(), password, time.Now()))

	return nil
}
