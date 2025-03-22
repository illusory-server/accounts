package entity

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/illusory-server/accounts/internal/domain"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/illusory-server/accounts/pkg/errors/codex"
	"github.com/illusory-server/accounts/pkg/errors/errx"
	"github.com/pkg/errors"
	"time"
)

const (
	MinNickLen = 2
	MaxNickLen = 124
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
	avatarLink domain.Option[vo.Link]
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
) (*Account, error) {
	result := &Account{
		id:         id,
		info:       info,
		role:       role,
		nickname:   nickname,
		password:   password,
		avatarLink: domain.NewEmptyOptional[vo.Link](),
		updatedAt:  updatedAt,
		createdAt:  createdAt,
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
	return a.id
}

func (a *Account) Info() vo.AccountInfo {
	return a.info
}

func (a *Account) Role() vo.Role {
	return a.role
}

func (a *Account) Nickname() string {
	return a.nickname
}

func (a *Account) Password() vo.Password {
	return a.password
}

func (a *Account) UpdatedAt() time.Time {
	return a.updatedAt
}

func (a *Account) CreatedAt() time.Time {
	return a.createdAt
}

func (a *Account) AvatarLink() domain.Option[vo.Link] {
	return a.avatarLink
}

// setters

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
	a.avatarLink = a.avatarLink.Set(link)
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
