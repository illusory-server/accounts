package factory

import (
	"github.com/illusory-server/accounts/internal/domain/aggregate"
	"github.com/illusory-server/accounts/internal/domain/entity"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/pkg/errors"
	"time"
)

//go:generate mockgen -package mock_factory -source account.go -destination ../../mock/app_factory/account.go

type (
	Timer interface {
		Now() time.Time
	}

	IDGenerator interface {
		GenerateID() string
	}

	AccountFactory interface {
		CreateAccount(
			firstName, lastName, email, nick, password string,
		) (*aggregate.Account, error)
	}

	AccountFactoryImpl struct {
		now   Timer
		genID IDGenerator
	}
)

func NewAccountFactory(timer Timer, generatorID IDGenerator) AccountFactoryImpl {
	return AccountFactoryImpl{
		now:   timer,
		genID: generatorID,
	}
}

func (a AccountFactoryImpl) CreateAccount(
	firstName, lastName, email, nick, password string,
) (*aggregate.Account, error) {
	id, err := vo.NewID(a.genID.GenerateID())
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
	t := a.now.Now()

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

	accAggregate, err := aggregate.NewAccount(acc)
	if err != nil {
		return nil, errors.Wrap(err, "[AccountFactory] aggregate.NewAccount")
	}

	return accAggregate, nil
}
