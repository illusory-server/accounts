package psql

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/illusory-server/accounts/internal/domain/aggregate"
	"github.com/illusory-server/accounts/internal/domain/repository"
	"github.com/illusory-server/accounts/pkg/logger"
)

var _ repository.AccountCommand = (*AccountCommand)(nil)

type AccountCommand struct {
	db  QueryExecutor
	log logger.Logger
}

func NewAccountCommand(log logger.Logger, pool QueryExecutor) (*AccountCommand, error) {
	res := &AccountCommand{
		db:  pool,
		log: log,
	}

	if err := res.Validate(); err != nil {
		return nil, err
	}

	return res, nil
}

func (a *AccountCommand) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.log, validation.Required),
	)
}

func (a *AccountCommand) Create(ctx context.Context, account *aggregate.Account) (*aggregate.Account, error) {
	return account, nil
}

func (a *AccountCommand) CreateMany(ctx context.Context, accounts []*aggregate.Account) error {
	return nil
}

func (a *AccountCommand) Update(ctx context.Context, account *aggregate.Account) error {
	return nil
}

func (a *AccountCommand) DeleteById(ctx context.Context, id string) error {
	return nil
}

func (a *AccountCommand) DeleteByEmail(ctx context.Context, email string) error {
	return nil
}

func (a *AccountCommand) DeleteByNickname(ctx context.Context, nickname string) error {
	return nil
}
