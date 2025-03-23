package account

import (
	"context"
	"github.com/illusory-server/accounts/pkg/errors/codex"
	"github.com/illusory-server/accounts/pkg/errors/errx"
	"github.com/illusory-server/accounts/pkg/logger"
	"github.com/pkg/errors"
)

func (a *AccountsUseCase) Create(ctx context.Context, firstName, lastName, email, nick, password string) (*WithoutPassword, error) {
	candidateByNick, err := a.accountQuery.HasByNickname(ctx, nick)
	if err != nil {
		a.log.Error(ctx, "has user by nickname query failed",
			logger.Err(err),
			logger.String("nickname", nick),
		)
		return nil, errors.Wrap(err, "[AccountUseCase] accountQuery.HasByNickname")
	}
	if candidateByNick {
		a.log.Debug(ctx, "has user by nickname query failed",
			logger.String("nickname", nick),
		)
		return nil, errx.New(codex.AlreadyExists, "account already exists")
	}
	candidateByEmail, err := a.accountQuery.HasByEmail(ctx, email)
	if err != nil {
		a.log.Error(ctx, "has user by email query failed",
			logger.Err(err),
			logger.String("email", email),
		)
	}
	if candidateByEmail {
		a.log.Debug(ctx, "has user by email query failed",
			logger.String("email", email),
		)
		return nil, errx.New(codex.AlreadyExists, "account email already exists")
	}

	acc, err := a.accountFactory.CreateAccount(firstName, lastName, email, nick, password)
	if err != nil {
		a.log.Error(ctx, "failed to create account",
			logger.Err(err),
			logger.String("first_name", firstName),
			logger.String("last_name", lastName),
			logger.String("email", email),
			logger.String("nickname", nick),
		)
		return nil, errors.Wrap(err, "[AccountUseCase] accountFactory.CreateAccount")
	}

	acc, err = a.accountCommand.Create(ctx, acc)
	if err != nil {
		a.log.Error(ctx, "failed to create account",
			logger.Err(err),
			logger.Any("account", acc),
		)
		return nil, errors.Wrap(err, "[AccountUseCase] accountCommand.Create")
	}

	return ConvertAccountAggregateToWithoutPassword(acc), nil
}
