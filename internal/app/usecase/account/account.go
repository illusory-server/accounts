package account

import (
	"context"
	"github.com/illusory-server/accounts/internal/app/factory"
	"github.com/illusory-server/accounts/internal/domain/entity"
	"github.com/illusory-server/accounts/internal/domain/repository"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/illusory-server/accounts/pkg/logger"
	"time"
)

type (
	WithoutPassword struct {
		Id        vo.ID
		FirstName string
		LastName  string
		Email     string
		Nickname  string
		Role      string
		UpdatedAt time.Time
		CreatedAt time.Time
	}

	// UseCase TODO - Требования по мульти create, update?
	UseCase interface {
		Create(ctx context.Context, firstName, lastName, email, nick, password string) (*WithoutPassword, error)

		UpdateById(ctx context.Context, id, firstName, lastName, nick string) (*WithoutPassword, error)
		UpdateEmailById(ctx context.Context, id, email string) (*WithoutPassword, error)
		UpdatePasswordById(ctx context.Context, id, oldPassword, password string) (*WithoutPassword, error)
		UpdateRoleById(ctx context.Context, id, role string) (*WithoutPassword, error)

		DeleteById(ctx context.Context, id string) error
		DeleteManyByIds(ctx context.Context, ids []string) error

		GetById(ctx context.Context, id string) (*WithoutPassword, error)
		GetWithPasswordById(ctx context.Context, id string) (*entity.Account, error)
		GetByEmail(ctx context.Context, email string) (*WithoutPassword, error)
		GetByNickname(ctx context.Context, nickname string) (*WithoutPassword, error)
		GetByQuery(ctx context.Context, query vo.Query) ([]*WithoutPassword, error)
		GetByIds(ctx context.Context, ids []string) ([]*WithoutPassword, error)
	}

	AccountsUseCase struct {
		log            logger.Logger
		accountFactory factory.AccountFactory
		accountQuery   repository.AccountQuery
		accountCommand repository.AccountCommand
	}
)

func (a *AccountsUseCase) GetWithPasswordById(ctx context.Context, id string) (*entity.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AccountsUseCase) DeleteById(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (a *AccountsUseCase) DeleteManyByIds(ctx context.Context, ids []string) error {
	//TODO implement me
	panic("implement me")
}

func (a *AccountsUseCase) GetById(ctx context.Context, id string) (*WithoutPassword, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AccountsUseCase) GetByEmail(ctx context.Context, email string) (*WithoutPassword, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AccountsUseCase) GetByNickname(ctx context.Context, nickname string) (*WithoutPassword, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AccountsUseCase) GetByQuery(ctx context.Context, query vo.Query) ([]*WithoutPassword, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AccountsUseCase) GetByIds(ctx context.Context, ids []string) ([]*WithoutPassword, error) {
	//TODO implement me
	panic("implement me")
}

func NewUseCase(
	accountFactory factory.AccountFactory,
) UseCase {
	return &AccountsUseCase{
		accountFactory: accountFactory,
	}
}

func NewWithoutPasswordFromEntity(entity *entity.Account) *WithoutPassword {
	return &WithoutPassword{
		Id:        entity.ID(),
		FirstName: entity.Info().FirstName(),
		LastName:  entity.Info().LastName(),
		Email:     entity.Info().Email(),
		Nickname:  entity.Nickname(),
		Role:      string(entity.Role().Value()),
		UpdatedAt: entity.UpdatedAt(),
		CreatedAt: entity.CreatedAt(),
	}
}
