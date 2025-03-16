package event

import (
	"github.com/illusory-server/accounts/internal/domain/vo"
	"time"
)

const (
	AccountChangeNicknameType Type = "NicknameChanged"
	AccountChangeInfoType     Type = "AccountChangeInfo"
	AccountChangePasswordType Type = "AccountChangePassword"
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
