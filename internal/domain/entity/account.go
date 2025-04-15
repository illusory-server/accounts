package entity

import (
	"encoding/json"
	"time"

	"github.com/illusory-server/accounts/pkg/fn"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/illusory-server/accounts/pkg/errors/codex"
	"github.com/illusory-server/accounts/pkg/errors/errx"
	"github.com/pkg/errors"
)

const (
	MinNickLen = 2
	MaxNickLen = 128
)

func validateTimeBeforeNow(value interface{}) error {
	t := value.(time.Time)
	if !t.Before(time.Now()) {
		return errors.New("invalid time value")
	}
	return nil
}

type Account struct {
	id         vo.ID
	info       vo.AccountInfo
	role       vo.Role
	nickname   string
	password   vo.Password
	avatarLink fn.Option[vo.Link]
	version    uint64
	updatedAt  time.Time
	createdAt  time.Time
}

func NewAccount(
	id vo.ID,
	info vo.AccountInfo,
	role vo.Role,
	nickname string,
	password vo.Password,
	updatedAt time.Time,
	createdAt time.Time,
	version uint64,
) (*Account, error) {
	result := &Account{
		id:         id,
		info:       info,
		role:       role,
		nickname:   nickname,
		password:   password,
		avatarLink: fn.None[vo.Link](),
		updatedAt:  updatedAt,
		createdAt:  createdAt,
		version:    version,
	}

	if err := result.Validate(); err != nil {
		return nil, errx.WrapWithCode(err, codex.InvalidArgument, "Account.Validate")
	}

	return result, nil
}

func (a *Account) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.id),
		validation.Field(&a.info),
		validation.Field(&a.role),
		validation.Field(&a.nickname, validation.Required, validation.Length(MinNickLen, MaxNickLen)),
		validation.Field(&a.password),
		validation.Field(&a.updatedAt, validation.By(validateTimeBeforeNow)),
		validation.Field(&a.createdAt, validation.By(validateTimeBeforeNow)),
	)
}

// getters

func (a *Account) ID() vo.ID {
	if a == nil {
		return vo.ID{}
	}
	return a.id
}

func (a *Account) Info() vo.AccountInfo {
	if a == nil {
		return vo.AccountInfo{}
	}
	return a.info
}

func (a *Account) Role() vo.Role {
	if a == nil {
		return vo.Role{}
	}
	return a.role
}

func (a *Account) Nickname() string {
	if a == nil {
		return ""
	}
	return a.nickname
}

func (a *Account) Password() vo.Password {
	if a == nil {
		return vo.Password{}
	}
	return a.password
}

func (a *Account) UpdatedAt() time.Time {
	if a == nil {
		return time.Time{}
	}
	return a.updatedAt
}

func (a *Account) CreatedAt() time.Time {
	if a == nil {
		return time.Time{}
	}
	return a.createdAt
}

func (a *Account) AvatarLink() fn.Option[vo.Link] {
	if a == nil {
		return fn.None[vo.Link]()
	}
	return a.avatarLink
}

func (a *Account) Version() uint64 {
	if a == nil {
		return 0
	}
	return a.version
}

// setters

func (a *Account) VersionIncrement() {
	if a == nil {
		return
	}
	a.version++
}

func (a *Account) SetInfo(info vo.AccountInfo) error {
	if err := info.Validate(); err != nil {
		return errx.WrapWithCode(err, codex.InvalidArgument, "AccountInfo.Validate")
	}
	a.info = info
	return nil
}

func (a *Account) SetNickname(nickname string) error {
	err := validation.Validate(nickname, validation.Required, validation.Length(MinNickLen, MaxNickLen))
	if err != nil {
		return errx.WrapWithCode(err, codex.InvalidArgument, "[Account] validation.Validate")
	}
	a.nickname = nickname
	return nil
}

func (a *Account) SetRole(role vo.Role) error {
	if err := role.Validate(); err != nil {
		return errx.WrapWithCode(err, codex.InvalidArgument, "[Account] validation.Validate")
	}
	a.role = role
	return nil
}

func (a *Account) SetPassword(password vo.Password) error {
	if err := password.Validate(); err != nil {
		return errx.WrapWithCode(err, codex.InvalidArgument, "[Account] validation.Validate")
	}
	a.password = password
	return nil
}

func (a *Account) SetAvatarLink(link vo.Link) error {
	if err := link.Validate(); err != nil {
		return errx.WrapWithCode(err, codex.InvalidArgument, "[Account] validation.Validate")
	}
	a.avatarLink = fn.Some(link)
	return nil
}

func (a *Account) SetUpdatedAt(updatedAt time.Time) error {
	err := validation.Validate(updatedAt, validation.By(validateTimeBeforeNow))
	if err != nil {
		return errx.WrapWithCode(err, codex.InvalidArgument, "[Account] validation.Validate")
	}
	a.updatedAt = updatedAt
	return nil
}

func (a *Account) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"id":          a.ID(),
		"info":        a.Info(),
		"role":        a.Role(),
		"nickname":    a.Nickname(),
		"avatar_link": a.AvatarLink().ValueOrDefault(vo.Link{}),
		"updated_at":  a.UpdatedAt(),
		"created_at":  a.CreatedAt(),
	}
	return json.Marshal(data)
}

type ReadOnlyAccount struct {
	acc *Account
}

func NewReadOnlyAccount(acc *Account) ReadOnlyAccount {
	return ReadOnlyAccount{acc: acc}
}

func (r ReadOnlyAccount) ID() vo.ID                      { return r.acc.ID() }
func (r ReadOnlyAccount) Info() vo.AccountInfo           { return r.acc.Info() }
func (r ReadOnlyAccount) Role() vo.Role                  { return r.acc.Role() }
func (r ReadOnlyAccount) Nickname() string               { return r.acc.Nickname() }
func (r ReadOnlyAccount) Password() vo.Password          { return r.acc.Password() }
func (r ReadOnlyAccount) AvatarLink() fn.Option[vo.Link] { return r.acc.AvatarLink() }
func (r ReadOnlyAccount) UpdatedAt() time.Time           { return r.acc.UpdatedAt() }
func (r ReadOnlyAccount) CreatedAt() time.Time           { return r.acc.CreatedAt() }
func (r ReadOnlyAccount) Version() uint64                { return r.acc.Version() }
