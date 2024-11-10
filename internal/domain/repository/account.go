package repository

import (
	"context"
	"github.com/illusory-server/accounts/internal/domain/model"
	"github.com/illusory-server/accounts/internal/domain/query"
	"github.com/illusory-server/accounts/internal/domain/vo"
)

type Accounts interface {
	Create(ctx context.Context, account *model.Account) (*model.Account, error)

	HasById(ctx context.Context, id vo.ID) (bool, error)
	HasByEmail(ctx context.Context, email string) (bool, error)
	HasByNickname(ctx context.Context, nickname string) (bool, error)

	GetById(ctx context.Context, id vo.ID) (*model.Account, error)
	GetByEmail(ctx context.Context, email string) (*model.Account, error)
	GetByNickname(ctx context.Context, nickname string) (*model.Account, error)
	GetByQuery(ctx context.Context, query *query.Pagination) ([]*model.Account, error)
	GetPageCountByLimit(ctx context.Context, limit uint64) (uint64, error)

	DeleteById(ctx context.Context, id vo.ID) error
	DeleteByEmail(ctx context.Context, email string) error
	DeleteByNickname(ctx context.Context, nickname string) error

	UpdateById(ctx context.Context, account *model.Account) (*model.Account, error)
	UpdatePasswordById(ctx context.Context, id vo.ID, newPassword string) (*model.Account, error)
	UpdateRoleById(ctx context.Context, id vo.ID, role string) (*model.Account, error)

	CheckAccountRoleById(ctx context.Context, id vo.ID, expectedRole string) (bool, error)
}
