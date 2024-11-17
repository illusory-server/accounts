package factory

import (
	"github.com/google/uuid"
	"github.com/illusory-server/accounts/internal/domain/entity"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/pkg/errors"
	"time"
)

type (
	AccountFactory interface {
		CreateAccount(
			firstName, lastName, email, nick, password string,
		) (*entity.Account, error)
	}

	AccountFactoryImpl struct{}
)

func NewAccountFactory() AccountFactoryImpl {
	return AccountFactoryImpl{}
}

func (a AccountFactoryImpl) CreateAccount(
	firstName, lastName, email, nick, password string,
) (*entity.Account, error) {
	id, err := vo.NewID(uuid.New().String())
	if err != nil {
		return nil, errors.Wrap(err, "[AccountFactory] vo.NewID")
	}
	info, err := vo.NewAccountInfo(firstName, lastName, email)
	if err != nil {
		return nil, errors.Wrap(err, "[AccountFactory] vo.NewAccountInfo")
	}
	role, err := vo.NewRole(vo.RoleUser)
	if err != nil {
		return nil, errors.Wrap(err, "[AccountFactory] vo.NewRole")
	}
	pass, err := vo.NewPassword(password)
	if err != nil {
		return nil, errors.Wrap(err, "[AccountFactory] vo.NewPassword")
	}
	t := time.Now()

	acc, err := entity.NewAccount(
		id,
		info,
		role,
		nick,
		pass,
		t,
		t,
	)
	if err != nil {
		return nil, errors.Wrap(err, "[AccountFactory] entity.NewAccount")
	}
	return acc, nil
}
