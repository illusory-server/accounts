package event

import (
	"time"

	"github.com/illusory-server/accounts/internal/domain/vo"
)

const (
	AccountChangeNicknameType   Type = "NicknameChanged"
	AccountChangeInfoType       Type = "AccountChangeInfo"
	AccountChangePasswordType   Type = "AccountChangePassword"
	AccountChangeEmailType      Type = "AccountChangeEmail"
	AccountChangeRoleType       Type = "AccountChangeRole"
	AccountChangeAvatarLinkType Type = "AccountChangeAvatarLink"
)

type AccountChangeNickname struct {
	id        vo.ID
	nickname  string
	timestamp time.Time
}

func (a AccountChangeNickname) ID() vo.ID            { return a.id }
func (a AccountChangeNickname) Nickname() string     { return a.nickname }
func (a AccountChangeNickname) Timestamp() time.Time { return a.timestamp }
func (a AccountChangeNickname) Type() Type           { return AccountChangeNicknameType }

func NewAccountChangeNickname(id vo.ID, nickname string, t time.Time) AccountChangeNickname {
	return AccountChangeNickname{
		id:        id,
		nickname:  nickname,
		timestamp: t,
	}
}

type AccountChangeInfo struct {
	id        vo.ID
	info      vo.AccountInfo
	timestamp time.Time
}

func (a AccountChangeInfo) ID() vo.ID            { return a.id }
func (a AccountChangeInfo) Info() vo.AccountInfo { return a.info }
func (a AccountChangeInfo) Timestamp() time.Time { return a.timestamp }
func (a AccountChangeInfo) Type() Type           { return AccountChangeInfoType }

func NewAccountChangeInfo(id vo.ID, info vo.AccountInfo, ts time.Time) AccountChangeInfo {
	return AccountChangeInfo{
		id:        id,
		info:      info,
		timestamp: ts,
	}
}

type AccountChangePassword struct {
	id        vo.ID
	password  vo.Password
	timestamp time.Time
}

func (a AccountChangePassword) ID() vo.ID             { return a.id }
func (a AccountChangePassword) Password() vo.Password { return a.password }
func (a AccountChangePassword) Timestamp() time.Time  { return a.timestamp }
func (a AccountChangePassword) Type() Type            { return AccountChangePasswordType }

func NewAccountChangePassword(id vo.ID, password vo.Password, ts time.Time) AccountChangePassword {
	return AccountChangePassword{
		id:        id,
		password:  password,
		timestamp: ts,
	}
}

type AccountChangeEmail struct {
	id        vo.ID
	email     string
	timestamp time.Time
}

func (a AccountChangeEmail) ID() vo.ID {
	return a.id
}

func (a AccountChangeEmail) Email() string {
	return a.email
}

func (a AccountChangeEmail) Timestamp() time.Time {
	return a.timestamp
}

func (a AccountChangeEmail) Type() Type {
	return AccountChangeEmailType
}

func NewAccountChangeEmail(id vo.ID, email string, t time.Time) AccountChangeEmail {
	return AccountChangeEmail{
		id:        id,
		email:     email,
		timestamp: t,
	}
}

type AccountChangeRole struct {
	id        vo.ID
	role      vo.Role
	timestamp time.Time
}

func (a AccountChangeRole) ID() vo.ID {
	return a.id
}

func (a AccountChangeRole) Role() vo.Role {
	return a.role
}

func (a AccountChangeRole) Timestamp() time.Time {
	return a.timestamp
}

func (a AccountChangeRole) Type() Type {
	return AccountChangeRoleType
}

func NewAccountChangeRole(id vo.ID, role vo.Role, ts time.Time) AccountChangeRole {
	return AccountChangeRole{
		id:        id,
		role:      role,
		timestamp: ts,
	}
}

type AccountChangeAvatarLink struct {
	id         vo.ID
	avatarLink vo.Link
	timestamp  time.Time
}

func (a AccountChangeAvatarLink) ID() vo.ID            { return a.id }
func (a AccountChangeAvatarLink) AvatarLink() vo.Link  { return a.avatarLink }
func (a AccountChangeAvatarLink) Timestamp() time.Time { return a.timestamp }
func (a AccountChangeAvatarLink) Type() Type           { return AccountChangeAvatarLinkType }

func NewAccountChangeAvatarLink(id vo.ID, avatarLink vo.Link, t time.Time) AccountChangeAvatarLink {
	return AccountChangeAvatarLink{
		id:         id,
		avatarLink: avatarLink,
		timestamp:  t,
	}
}
