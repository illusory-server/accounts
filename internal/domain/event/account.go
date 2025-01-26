package event

import (
	"github.com/illusory-server/accounts/internal/domain/vo"
	"time"
)

type AccountChangeNickname struct {
	ID       vo.ID
	Nickname string
	Time     time.Time
}

func (a AccountChangeNickname) Type() Type {
	return "AccountChangeNickname"
}

func (a AccountChangeNickname) Timestamp() time.Time {
	return a.Time
}

type AccountChangeInfo struct {
	ID   vo.ID
	Info vo.AccountInfo
	Time time.Time
}

func (a AccountChangeInfo) Type() Type {
	return "AccountChangeInfo"
}

func (a AccountChangeInfo) Timestamp() time.Time {
	return a.Time
}

type AccountChangePassword struct {
	ID       vo.ID
	Password vo.Password
	Time     time.Time
}

func (a AccountChangePassword) Type() Type {
	return "AccountChangePassword"
}

func (a AccountChangePassword) Timestamp() time.Time {
	return a.Time
}
