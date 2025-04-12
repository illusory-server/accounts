package psql

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/illusory-server/accounts/internal/app/factory"
	"github.com/illusory-server/accounts/internal/domain/aggregate"
	"github.com/illusory-server/accounts/internal/domain/repository"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/illusory-server/accounts/pkg/logger"
	"github.com/pkg/errors"
	"time"
)

type timer struct{}

func (t timer) Now() time.Time {
	return time.Now()
}

type idGen struct{}

func (i idGen) GenerateID() string {
	return uuid.New().String()
}

var _ repository.AccountQuery = (*AccountQuery)(nil)

type AccountQuery struct {
	db  QueryExecutor
	log logger.Logger
}

func NewAccountQuery(log logger.Logger, pool QueryExecutor) (*AccountQuery, error) {
	res := &AccountQuery{
		db:  pool,
		log: log,
	}

	if err := res.Validate(); err != nil {
		return nil, err
	}

	return res, nil
}

func (a *AccountQuery) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.log, validation.Required),
	)
}

func (a *AccountQuery) HasById(ctx context.Context, id string) (bool, error) {
	return true, nil
}

func (a *AccountQuery) HasByEmail(ctx context.Context, email string) (bool, error) {
	return true, nil
}

func (a *AccountQuery) HasByNickname(ctx context.Context, nickname string) (bool, error) {
	return true, nil
}

func (a *AccountQuery) GetById(ctx context.Context, id string) (*aggregate.Account, error) {
	accFactory := factory.NewAccountFactory(timer{}, idGen{})
	aggr, err := accFactory.CreateAccount("marlen", "karimov", "haha@gmail.com", "eer0", "Crutoi123456@")
	if err != nil {
		return nil, errors.Wrap(err, "[AccountQuery] accFactory.CreateAccount")
	}
	return aggr, nil
}

func (a *AccountQuery) GetByIds(ctx context.Context, ids []string) ([]*aggregate.Account, error) {
	accFactory := factory.NewAccountFactory(timer{}, idGen{})
	aggr, err := accFactory.CreateAccount("marlen", "karimov", "haha@gmail.com", "eer0", "Crutoi123456@")
	if err != nil {
		return nil, errors.Wrap(err, "[AccountQuery] accFactory.CreateAccount")
	}
	return []*aggregate.Account{aggr}, nil
}

func (a *AccountQuery) GetByEmail(ctx context.Context, email string) (*aggregate.Account, error) {
	accFactory := factory.NewAccountFactory(timer{}, idGen{})
	aggr, err := accFactory.CreateAccount("marlen", "karimov", "haha@gmail.com", "eer0", "Crutoi123456@")
	if err != nil {
		return nil, errors.Wrap(err, "[AccountQuery] accFactory.CreateAccount")
	}
	return aggr, nil
}

func (a *AccountQuery) GetByNickname(ctx context.Context, nickname string) (*aggregate.Account, error) {
	accFactory := factory.NewAccountFactory(timer{}, idGen{})
	aggr, err := accFactory.CreateAccount("marlen", "karimov", "haha@gmail.com", "eer0", "Crutoi123456@")
	if err != nil {
		return nil, errors.Wrap(err, "[AccountQuery] accFactory.CreateAccount")
	}
	return aggr, nil
}

func (a *AccountQuery) GetByQuery(ctx context.Context, query vo.Query) ([]*aggregate.Account, uint, error) {
	accFactory := factory.NewAccountFactory(timer{}, idGen{})
	aggr, err := accFactory.CreateAccount("marlen", "karimov", "haha@gmail.com", "eer0", "Crutoi123456@")
	if err != nil {
		return nil, 0, errors.Wrap(err, "[AccountQuery] accFactory.CreateAccount")
	}
	return []*aggregate.Account{aggr}, 24, nil //nolint:mnd
}

func (a *AccountQuery) GetPageCountByLimit(ctx context.Context, limit uint64) (uint64, error) {
	return 100, nil //nolint:mnd
}

func (a *AccountQuery) CheckAccountRoleById(
	ctx context.Context, id string, expectedRole vo.AccountRoleType,
) (bool, error) {
	return true, nil
}
