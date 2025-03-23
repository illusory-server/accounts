package aggregate

import (
	"encoding/json"
	"github.com/illusory-server/accounts/internal/domain"
	"github.com/illusory-server/accounts/internal/domain/entity"
	"github.com/illusory-server/accounts/internal/domain/event"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/illusory-server/accounts/pkg/errors/codex"
	"github.com/illusory-server/accounts/pkg/errors/errx"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type ReadOnlyAccountEntity struct {
	acc *entity.Account
}

func NewReadOnlyAccountEntity(acc *entity.Account) ReadOnlyAccountEntity {
	return ReadOnlyAccountEntity{acc: acc}
}

func (r ReadOnlyAccountEntity) ID() vo.ID                          { return r.acc.ID() }
func (r ReadOnlyAccountEntity) Info() vo.AccountInfo               { return r.acc.Info() }
func (r ReadOnlyAccountEntity) Role() vo.Role                      { return r.acc.Role() }
func (r ReadOnlyAccountEntity) Nickname() string                   { return r.acc.Nickname() }
func (r ReadOnlyAccountEntity) Password() vo.Password              { return r.acc.Password() }
func (r ReadOnlyAccountEntity) AvatarLink() domain.Option[vo.Link] { return r.acc.AvatarLink() }
func (r ReadOnlyAccountEntity) UpdatedAt() time.Time               { return r.acc.UpdatedAt() }
func (r ReadOnlyAccountEntity) CreatedAt() time.Time               { return r.acc.CreatedAt() }

type Account struct {
	account *entity.Account
	events  event.Events
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

func (a *Account) Account() ReadOnlyAccountEntity {
	return NewReadOnlyAccountEntity(a.account)
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

func (a *Account) ChangeNickname(nickname string, t time.Time) error {
	err := a.account.SetNickname(nickname)
	if err != nil {
		return errors.Wrap(err, "[Account] account.SetNickname")
	}

	a.AddEvent(event.NewAccountChangeNickname(a.account.ID(), nickname, t))

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

func (a *Account) ChangeAccountInfo(info vo.AccountInfo, t time.Time) error {
	err := a.account.SetInfo(info)
	if err != nil {
		return errors.Wrap(err, "[Account] account.SetInfo")
	}

	a.AddEvent(event.NewAccountChangeInfo(a.account.ID(), info, t))

	return nil
}

func (a *Account) ChangeEmail(info vo.AccountInfo, t time.Time) error {
	err := a.account.SetInfo(info)
	if err != nil {
		return errors.Wrap(err, "[Account] account.SetInfo")
	}

	a.AddEvent(event.NewAccountChangeEmail(a.account.ID(), info.Email(), t))

	return nil
}

func (a *Account) ChangeRole(role vo.Role, t time.Time) error {
	err := a.account.SetRole(role)
	if err != nil {
		return errors.Wrap(err, "[Account] account.SetRole")
	}

	a.AddEvent(event.NewAccountChangeRole(a.account.ID(), role, t))

	return nil
}

func (a *Account) ChangeAvatarLink(link vo.Link, t time.Time) error {
	err := a.account.SetAvatarLink(link)
	if err != nil {
		return errors.Wrap(err, "[Account] account.SetAvatarLink")
	}

	a.AddEvent(event.NewAccountChangeAvatarLink(a.account.ID(), link, t))

	return nil
}

func (a *Account) ChangePassword(password vo.Password, t time.Time) error {
	err := a.account.SetPassword(password)
	if err != nil {
		return errors.Wrap(err, "[Account] account.SetPassword")
	}

	a.AddEvent(event.NewAccountChangePassword(a.account.ID(), password, t))

	return nil
}

func (a *Account) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"account": a.account,
		"events":  a.events,
	}
	return json.Marshal(data)
}
