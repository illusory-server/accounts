package account

import (
	"context"
	"github.com/illusory-server/accounts/pkg/errors/codes"
	"github.com/illusory-server/accounts/pkg/errors/xerr"
	"github.com/illusory-server/accounts/pkg/logger/log"
)

var (
	ErrNicknameExist = xerr.New(codes.Conflict, "account nickname already exists")
	ErrEmailExist    = xerr.New(codes.Conflict, "account email already exists")
)

func (a *AccountsUseCase) Create(ctx context.Context, firstName, lastName, email, nick, password string) (*WithoutPassword, error) {
	candidateByNick, err := a.accountQuery.HasByNickname(ctx, nick)
	if err != nil {
		a.log.Error(ctx, "has user by nickname query failed",
			log.Err(err),
			log.String("nickname", nick),
		)
		return nil, xerr.Wrap(err, "[AccountUseCase] accountQuery.HasByNickname")
	}
	if candidateByNick {
		a.log.Debug(ctx, "has user by nickname query failed",
			log.String("nickname", nick),
		)
		return nil, ErrNicknameExist
	}
	candidateByEmail, err := a.accountQuery.HasByEmail(ctx, email)
	if err != nil {
		a.log.Error(ctx, "has user by email query failed",
			log.Err(err),
			log.String("email", email),
		)
	}
	if candidateByEmail {
		a.log.Debug(ctx, "has user by email query failed",
			log.String("email", email),
		)
		return nil, ErrEmailExist
	}

	acc, err := a.accountFactory.CreateAccount(firstName, lastName, email, nick, password)
	if err != nil {
		a.log.Error(ctx, "failed to create account",
			log.Err(err),
			log.String("first_name", firstName),
			log.String("last_name", lastName),
			log.String("email", email),
			log.String("nickname", nick),
		)
		return nil, xerr.Wrap(err, "[AccountUseCase] accountFactory.CreateAccount")
	}

	acc, err = a.accountCommand.Create(ctx, acc)
	if err != nil {
		a.log.Error(ctx, "failed to create account",
			log.Err(err),
			log.Any("account", acc),
		)
		return nil, xerr.Wrap(err, "[AccountUseCase] accountCommand.Create")
	}

	return NewWithoutPasswordFromEntity(acc), nil
}
