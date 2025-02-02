package account

import (
	"context"
	"github.com/illusory-server/accounts/internal/domain/aggregate"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/illusory-server/accounts/pkg/logger/log"
	"github.com/pkg/errors"
	"strings"
)

func (a *AccountsUseCase) GetWithPasswordById(ctx context.Context, id string) (*aggregate.Account, error) {
	result, err := a.accountQuery.GetById(ctx, id)
	if err != nil {
		a.log.Error(ctx, "failed get account by id",
			log.String("id", id))
		return nil, errors.Wrap(err, "[AccountsUseCase] accountQuery.GetById")
	}
	return result, nil
}

func (a *AccountsUseCase) GetById(ctx context.Context, id string) (*WithoutPassword, error) {
	result, err := a.accountQuery.GetById(ctx, id)
	if err != nil {
		a.log.Error(ctx, "failed get account by id",
			log.String("id", id))
		return nil, errors.Wrap(err, "[AccountsUseCase] accountQuery.GetById")
	}
	return NewWithoutPasswordFromAggregate(result), nil
}

func (a *AccountsUseCase) GetByEmail(ctx context.Context, email string) (*WithoutPassword, error) {
	result, err := a.accountQuery.GetByEmail(ctx, email)
	if err != nil {
		a.log.Error(ctx, "failed get account by email",
			log.String("email", email))
		return nil, errors.Wrap(err, "[AccountsUseCase] accountQuery.GetByEmail")
	}
	return NewWithoutPasswordFromAggregate(result), nil
}

func (a *AccountsUseCase) GetByNickname(ctx context.Context, nickname string) (*WithoutPassword, error) {
	result, err := a.accountQuery.GetByNickname(ctx, nickname)
	if err != nil {
		a.log.Error(ctx, "failed get account by nickname",
			log.String("nickname", nickname))
		return nil, errors.Wrap(err, "[AccountsUseCase] accountQuery.GetByNickname")
	}
	return NewWithoutPasswordFromAggregate(result), nil
}

func (a *AccountsUseCase) GetByQuery(ctx context.Context, query vo.Query) ([]*WithoutPassword, error) {
	result, err := a.accountQuery.GetByQuery(ctx, query)
	if err != nil {
		a.log.Error(ctx, "failed get account by query",
			log.Any("query", query))
		return nil, errors.Wrap(err, "[AccountsUseCase] accountQuery.GetByQuery")
	}
	return NewWithoutPasswordsFromAggregates(result), nil
}

func (a *AccountsUseCase) GetByIds(ctx context.Context, ids []string) ([]*WithoutPassword, error) {
	result, err := a.accountQuery.GetByIds(ctx, ids)
	if err != nil {
		a.log.Error(ctx, "failed get account by ids",
			log.String("ids", strings.Join(ids, ",")))
		return nil, errors.Wrap(err, "[AccountsUseCase] accountQuery.GetByIds")
	}

	return NewWithoutPasswordsFromAggregates(result), nil
}
