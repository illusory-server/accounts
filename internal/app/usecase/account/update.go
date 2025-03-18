package account

import (
	"context"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/illusory-server/accounts/pkg/errors/codes"
	"github.com/illusory-server/accounts/pkg/errors/errx"
	"github.com/pkg/errors"
)

var (
	ErrOldPasswordNotEqual = errx.New(codes.PermissionDenied, "old password not equal")
)

func (a *AccountsUseCase) UpdateInfoById(ctx context.Context, id, firstName, lastName string) error {
	aggregate, err := a.accountQuery.GetById(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] accountQuery.GetById")
	}
	newInfo, err := vo.NewAccountInfo(firstName, lastName, aggregate.Account().Info().Email())
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] vo.NewAccountInfo")
	}
	err = aggregate.ChangeAccountInfo(newInfo)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] aggregate.ChangeAccountInfo")
	}
	err = a.accountCommand.Update(ctx, aggregate)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] a.accountCommand.Update")
	}
	return nil
}

func (a *AccountsUseCase) UpdateEmailById(ctx context.Context, id, email string) error {
	aggregate, err := a.accountQuery.GetById(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] accountQuery.GetById")
	}
	newEmail, err := vo.NewAccountInfo(
		aggregate.Account().Info().FirstName(),
		aggregate.Account().Info().LastName(),
		email,
	)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] vo.NewAccountInfo")
	}
	err = aggregate.ChangeEmail(newEmail)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] aggregate.ChangeEmail")
	}
	err = a.accountCommand.Update(ctx, aggregate)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] a.accountCommand.Update")
	}
	return nil
}

func (a *AccountsUseCase) UpdatePasswordById(ctx context.Context, id, oldPassword, password string) error {
	aggregate, err := a.accountQuery.GetById(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] accountQuery.GetById")
	}
	err = aggregate.ComparePassword(oldPassword)
	if err != nil {
		return errors.Wrap(ErrOldPasswordNotEqual, "[AccountsUseCase] aggregate.ComparePassword")
	}
	newPass, err := vo.NewPassword(password)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] vo.NewPassword")
	}
	err = aggregate.ChangePassword(newPass)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] aggregate.ChangePassword")
	}
	err = a.accountCommand.Update(ctx, aggregate)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] a.accountCommand.Update")
	}
	return nil
}

func (a *AccountsUseCase) UpdateRoleById(ctx context.Context, id, role string) error {
	aggregate, err := a.accountQuery.GetById(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] accountQuery.GetById")
	}
	newRole, err := vo.NewRole(vo.AccountRoleType(role))
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] vo.NewRole")
	}
	err = aggregate.ChangeRole(newRole)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] aggregate.ChangeRole")
	}
	err = a.accountCommand.Update(ctx, aggregate)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] a.accountCommand.Update")
	}
	return nil
}

func (a *AccountsUseCase) UpdateNickById(ctx context.Context, id, nick string) error {
	aggregate, err := a.accountQuery.GetById(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] accountQuery.GetById")
	}
	err = aggregate.ChangeNickname(nick)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] vo.NewRole")
	}
	err = a.accountCommand.Update(ctx, aggregate)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] a.accountCommand.Update")
	}
	return nil
}

func (a *AccountsUseCase) AddAvatarLink(ctx context.Context, id, url string) error {
	aggregate, err := a.accountQuery.GetById(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] accountQuery.GetById")
	}
	link, err := vo.NewLink(url)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] vo.NewLink")
	}
	err = aggregate.ChangeAvatarLink(link)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] aggregate.ChangeAvatarLink")
	}
	err = a.accountCommand.Update(ctx, aggregate)
	if err != nil {
		return errors.Wrap(err, "[AccountsUseCase] accountCommand.Update")
	}
	return nil
}
