package account

import (
	"context"
	"github.com/illusory-server/accounts/internal/app/factory"
	"github.com/illusory-server/accounts/internal/domain/aggregate"
	"github.com/illusory-server/accounts/internal/domain/repository"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/illusory-server/accounts/pkg/logger"
	"time"
)

//go:generate mockgen -package mock_usecase -source account.go -destination ../../../mock/usecase/account.go

type (
	WithoutPassword struct {
		ID        vo.ID
		FirstName string
		LastName  string
		Email     string
		Nickname  string
		Role      string
		AvatarURL string
		UpdatedAt time.Time
		CreatedAt time.Time
	}

	// UseCase TODO - Требования по мульти create, update?
	UseCase interface {
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

func NewUseCase(
	log logger.Logger,
	accountFactory factory.AccountFactory,
	accountQuery repository.AccountQuery,
	accountCommand repository.AccountCommand,
) *AccountsUseCase {
	return &AccountsUseCase{
		log:            log,
		accountFactory: accountFactory,
		accountQuery:   accountQuery,
		accountCommand: accountCommand,
	}
}

func NewWithoutPasswordFromAggregate(aggregate *aggregate.Account) *WithoutPassword {
	return &WithoutPassword{
		ID:        aggregate.Account().ID(),
		FirstName: aggregate.Account().Info().FirstName(),
		LastName:  aggregate.Account().Info().LastName(),
		Email:     aggregate.Account().Info().Email(),
		Nickname:  aggregate.Account().Nickname(),
		Role:      string(aggregate.Account().Role().Value()),
		AvatarURL: aggregate.Account().AvatarLink().ValueOrDefault(vo.Link{}).Value(),
		UpdatedAt: aggregate.Account().UpdatedAt(),
		CreatedAt: aggregate.Account().CreatedAt(),
	}
}

func NewWithoutPasswordsFromAggregates(aggregates []*aggregate.Account) []*WithoutPassword {
	result := make([]*WithoutPassword, 0, len(aggregates))
	for _, acc := range aggregates {
		result = append(result, NewWithoutPasswordFromAggregate(acc))
	}
	return result
}
