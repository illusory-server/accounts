package account

import (
	"context"
	"time"

	"github.com/illusory-server/accounts/internal/app/factory"
	"github.com/illusory-server/accounts/internal/domain/aggregate"
	"github.com/illusory-server/accounts/internal/domain/repository"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/illusory-server/accounts/pkg/logger"
)

var _ Account = (*UseCase)(nil)

//go:generate mockgen -package mock_usecase -source account.go -destination ../../../mock/usecase/account.go

type (
	Timer interface {
		Now() time.Time
	}

	// Account TODO - Требования по мульти create, update?
	Account interface {
		Create(ctx context.Context, firstName, lastName, email, nick, password string) (*WithoutPassword, error)

		UpdateInfoById(ctx context.Context, id, firstName, lastName string) error
		UpdateNickById(ctx context.Context, id, nick string) error
		UpdateEmailById(ctx context.Context, id, email string) error
		UpdatePasswordById(ctx context.Context, id, oldPassword, password string) error
		UpdateRoleById(ctx context.Context, id, role string) error
		AddAvatarLink(ctx context.Context, id, url string) error

		DeleteById(ctx context.Context, id string) error
		DeleteManyByIds(ctx context.Context, ids []string) error

		GetById(ctx context.Context, id string) (*WithoutPassword, error)
		GetWithPasswordById(ctx context.Context, id string) (*aggregate.Account, error)
		GetByEmail(ctx context.Context, email string) (*WithoutPassword, error)
		GetByNickname(ctx context.Context, nickname string) (*WithoutPassword, error)
		GetByQuery(ctx context.Context, query vo.Query) ([]*WithoutPassword, uint, error)
		GetByIds(ctx context.Context, ids []string) ([]*WithoutPassword, error)
	}

	UseCase struct {
		log            logger.Logger
		accountFactory factory.AccountFactory
		accountQuery   repository.AccountQuery
		accountCommand repository.AccountCommand
		now            Timer
	}
)

func NewUseCase(
	log logger.Logger,
	accountFactory factory.AccountFactory,
	accountQuery repository.AccountQuery,
	accountCommand repository.AccountCommand,
	nowTimer Timer,
) (*UseCase, error) {
	return &UseCase{
		log:            log,
		accountFactory: accountFactory,
		accountQuery:   accountQuery,
		accountCommand: accountCommand,
		now:            nowTimer,
	}, nil
}

func ConvertAccountAggregateToWithoutPassword(aggregate *aggregate.Account) *WithoutPassword {
	return &WithoutPassword{
		id:        aggregate.Account().ID(),
		firstName: aggregate.Account().Info().FirstName(),
		lastName:  aggregate.Account().Info().LastName(),
		email:     aggregate.Account().Info().Email(),
		nickname:  aggregate.Account().Nickname(),
		role:      string(aggregate.Account().Role().Value()),
		avatarURL: aggregate.Account().AvatarLink().ValueOrDefault(vo.Link{}).Value(),
		updatedAt: aggregate.Account().UpdatedAt(),
		createdAt: aggregate.Account().CreatedAt(),
	}
}
