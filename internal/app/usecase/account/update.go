package account

import (
	"context"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/illusory-server/accounts/pkg/errors/codex"
	"github.com/illusory-server/accounts/pkg/errors/errx"
	"github.com/illusory-server/accounts/pkg/logger"
	"github.com/pkg/errors"
)

var (
	ErrOldPasswordNotEqual = errx.New(codex.InvalidArgument, "old password not equal")
)

func (a *AccountsUseCase) UpdateInfoById(ctx context.Context, id, firstName, lastName string) error {
	aggregate, err := a.accountQuery.GetById(ctx, id)
	if err != nil {
		a.log.Error(
			ctx,
			"query get account by id failed",
			logger.String("id", id),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] accountQuery.GetById")
	}
	newInfo, err := vo.NewAccountInfo(firstName, lastName, aggregate.Account().Info().Email())
	if err != nil {
		a.log.Error(
			ctx,
			"new account info values object create failed",
			logger.String("first_name", firstName),
			logger.String("last_name", lastName),
			logger.String("email", aggregate.Account().Info().Email()),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] vo.NewAccountInfo")
	}
	err = aggregate.ChangeAccountInfo(newInfo, a.now.Now())
	if err != nil {
		a.log.Error(
			ctx,
			"aggregate change account info failed",
			logger.Any("info", newInfo),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] aggregate.ChangeAccountInfo")
	}
	err = a.accountCommand.Update(ctx, aggregate)
	if err != nil {
		a.log.Error(
			ctx,
			"command update account aggregate failed",
			logger.Any("aggregate", aggregate),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] accountCommand.Update")
	}
	return nil
}

func (a *AccountsUseCase) UpdateEmailById(ctx context.Context, id, email string) error {
	aggregate, err := a.accountQuery.GetById(ctx, id)
	if err != nil {
		a.log.Error(
			ctx,
			"query get account by id failed",
			logger.String("id", id),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] accountQuery.GetById")
	}
	newEmail, err := vo.NewAccountInfo(
		aggregate.Account().Info().FirstName(),
		aggregate.Account().Info().LastName(),
		email,
	)
	if err != nil {
		a.log.Error(
			ctx,
			"new account info values object create failed",
			logger.String("first_name", aggregate.Account().Info().FirstName()),
			logger.String("last_name", aggregate.Account().Info().LastName()),
			logger.String("email", email),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] vo.NewAccountInfo")
	}
	err = aggregate.ChangeEmail(newEmail, a.now.Now())
	if err != nil {
		a.log.Error(
			ctx,
			"aggregate change email failed",
			logger.Any("info", newEmail),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] aggregate.ChangeEmail")
	}
	err = a.accountCommand.Update(ctx, aggregate)
	if err != nil {
		a.log.Error(
			ctx,
			"command update account aggregate failed",
			logger.Any("aggregate", aggregate),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] a.accountCommand.Update")
	}
	return nil
}

func (a *AccountsUseCase) UpdatePasswordById(ctx context.Context, id, oldPassword, password string) error {
	aggregate, err := a.accountQuery.GetById(ctx, id)
	if err != nil {
		a.log.Error(
			ctx,
			"query get account by id failed",
			logger.String("id", id),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] accountQuery.GetById")
	}
	err = aggregate.ComparePassword(oldPassword)
	if err != nil {
		a.log.Info(
			ctx,
			"old password not equal",
			logger.String("id", id),
			logger.Err(err),
		)
		return errors.Wrap(ErrOldPasswordNotEqual, "[AccountsUseCase] aggregate.ComparePassword")
	}
	newPass, err := vo.NewPassword(password)
	if err != nil {
		a.log.Error(
			ctx,
			"new password values object create failed",
			logger.String("id", id),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] vo.NewPassword")
	}
	err = aggregate.ChangePassword(newPass, a.now.Now())
	if err != nil {
		a.log.Error(
			ctx,
			"aggregate change password failed",
			logger.Any("aggregate", aggregate),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] aggregate.ChangePassword")
	}
	err = a.accountCommand.Update(ctx, aggregate)
	if err != nil {
		a.log.Error(
			ctx,
			"command update account aggregate failed",
			logger.Any("aggregate", aggregate),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] a.accountCommand.Update")
	}
	return nil
}

func (a *AccountsUseCase) UpdateRoleById(ctx context.Context, id, role string) error {
	aggregate, err := a.accountQuery.GetById(ctx, id)
	if err != nil {
		a.log.Error(
			ctx,
			"query get account by id failed",
			logger.String("id", id),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] accountQuery.GetById")
	}
	newRole, err := vo.NewRole(vo.AccountRoleType(role))
	if err != nil {
		a.log.Error(
			ctx,
			"new role values object create failed",
			logger.String("role", role),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] vo.NewRole")
	}
	err = aggregate.ChangeRole(newRole, a.now.Now())
	if err != nil {
		a.log.Error(
			ctx,
			"aggregate change role failed",
			logger.Any("role", newRole),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] aggregate.ChangeRole")
	}
	err = a.accountCommand.Update(ctx, aggregate)
	if err != nil {
		a.log.Error(
			ctx,
			"command update account aggregate failed",
			logger.Any("aggregate", aggregate),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] a.accountCommand.Update")
	}
	return nil
}

func (a *AccountsUseCase) UpdateNickById(ctx context.Context, id, nick string) error {
	aggregate, err := a.accountQuery.GetById(ctx, id)
	if err != nil {
		a.log.Error(
			ctx,
			"query get account by id failed",
			logger.String("id", id),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] accountQuery.GetById")
	}
	err = aggregate.ChangeNickname(nick, a.now.Now())
	if err != nil {
		a.log.Error(
			ctx,
			"aggregate change nickname failed",
			logger.String("nickname", nick),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] vo.NewRole")
	}
	err = a.accountCommand.Update(ctx, aggregate)
	if err != nil {
		a.log.Error(
			ctx,
			"command update account aggregate failed",
			logger.Any("aggregate", aggregate),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] a.accountCommand.Update")
	}
	return nil
}

func (a *AccountsUseCase) AddAvatarLink(ctx context.Context, id, url string) error {
	aggregate, err := a.accountQuery.GetById(ctx, id)
	if err != nil {
		a.log.Error(
			ctx,
			"query get account by id failed",
			logger.String("id", id),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] accountQuery.GetById")
	}
	link, err := vo.NewLink(url)
	if err != nil {
		a.log.Error(
			ctx,
			"new link values object create failed",
			logger.String("url", url),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] vo.NewLink")
	}
	err = aggregate.ChangeAvatarLink(link, a.now.Now())
	if err != nil {
		a.log.Error(
			ctx,
			"aggregate change avatar link failed",
			logger.Any("link", link),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] aggregate.ChangeAvatarLink")
	}
	err = a.accountCommand.Update(ctx, aggregate)
	if err != nil {
		a.log.Error(
			ctx,
			"command update account aggregate failed",
			logger.Any("aggregate", aggregate),
			logger.Err(err),
		)
		return errors.Wrap(err, "[AccountsUseCase] accountCommand.Update")
	}
	return nil
}
