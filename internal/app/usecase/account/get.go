package account

import (
	"context"
	"strings"

	"github.com/illusory-server/accounts/internal/domain/aggregate"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/illusory-server/accounts/pkg/fn"
	"github.com/illusory-server/accounts/pkg/logger"
	"github.com/pkg/errors"
)

func (a *AccountsUseCase) GetWithPasswordById(ctx context.Context, id string) (*aggregate.Account, error) {
	result, err := a.accountQuery.GetById(ctx, id)
	if err != nil {
		a.log.Error(ctx, "failed get account by id",
			logger.String("id", id))
		return nil, errors.Wrap(err, "[AccountsUseCase] accountQuery.GetById")
	}
	return result, nil
}

func (a *AccountsUseCase) GetById(ctx context.Context, id string) (*WithoutPassword, error) {
	result, err := a.accountQuery.GetById(ctx, id)
	if err != nil {
		a.log.Error(ctx, "failed get account by id",
			logger.String("id", id))
		return nil, errors.Wrap(err, "[AccountsUseCase] accountQuery.GetById")
	}
	return ConvertAccountAggregateToWithoutPassword(result), nil
}

func (a *AccountsUseCase) GetByEmail(ctx context.Context, email string) (*WithoutPassword, error) {
	result, err := a.accountQuery.GetByEmail(ctx, email)
	if err != nil {
		a.log.Error(ctx, "failed get account by email",
			logger.String("email", email))
		return nil, errors.Wrap(err, "[AccountsUseCase] accountQuery.GetByEmail")
	}
	return ConvertAccountAggregateToWithoutPassword(result), nil
}

func (a *AccountsUseCase) GetByNickname(ctx context.Context, nickname string) (*WithoutPassword, error) {
	result, err := a.accountQuery.GetByNickname(ctx, nickname)
	if err != nil {
		a.log.Error(ctx, "failed get account by nickname",
			logger.String("nickname", nickname))
		return nil, errors.Wrap(err, "[AccountsUseCase] accountQuery.GetByNickname")
	}
	return ConvertAccountAggregateToWithoutPassword(result), nil
}

func (a *AccountsUseCase) GetByQuery(ctx context.Context, query vo.Query) ([]*WithoutPassword, uint, error) {
	result, pageCount, err := a.accountQuery.GetByQuery(ctx, query)
	if err != nil {
		a.log.Error(ctx, "failed get account by query",
			logger.Any("query", query))
		return nil, 0, errors.Wrap(err, "[AccountsUseCase] accountQuery.GetByQuery")
	}
	return fn.Map(result, ConvertAccountAggregateToWithoutPassword), pageCount, nil
}

func (a *AccountsUseCase) GetByIds(ctx context.Context, ids []string) ([]*WithoutPassword, error) {
	result, err := a.accountQuery.GetByIds(ctx, ids)
	if err != nil {
		a.log.Error(ctx, "failed get account by ids",
			logger.String("ids", strings.Join(ids, ",")))
		return nil, errors.Wrap(err, "[AccountsUseCase] accountQuery.GetByIds")
	}

	return fn.Map(result, ConvertAccountAggregateToWithoutPassword), nil
}
