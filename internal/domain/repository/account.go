package repository

import (
	"context"
	"github.com/illusory-server/accounts/internal/domain/aggregate"

	"github.com/illusory-server/accounts/internal/domain/vo"
)

type AccountCommand interface {
	Create(ctx context.Context, account *aggregate.Account) (*aggregate.Account, error)
	CreateMany(ctx context.Context, accounts []*aggregate.Account) error

	UpdateInfoById(ctx context.Context, id vo.ID, info vo.AccountInfo) error
	UpdateNicknameById(ctx context.Context, id vo.ID, nickname string) error
	UpdateAvatarLinkById(ctx context.Context, id vo.ID, link vo.Link) error
	UpdatePasswordById(ctx context.Context, id vo.ID, newPassword vo.Password) error
	UpdateRoleById(ctx context.Context, id vo.ID, role vo.Role) error

	DeleteById(ctx context.Context, id string) error
	DeleteByEmail(ctx context.Context, email string) error
	DeleteByNickname(ctx context.Context, nickname string) error
}

type AccountQuery interface {
	HasById(ctx context.Context, id string) (bool, error)
	HasByEmail(ctx context.Context, email string) (bool, error)
	HasByNickname(ctx context.Context, nickname string) (bool, error)

	GetById(ctx context.Context, id string) (*aggregate.Account, error)
	GetByIds(ctx context.Context, ids []string) ([]*aggregate.Account, error)
	GetByEmail(ctx context.Context, email string) (*aggregate.Account, error)
	GetByNickname(ctx context.Context, nickname string) (*aggregate.Account, error)
	GetByQuery(ctx context.Context, query vo.Query) ([]*aggregate.Account, error)

	GetPageCountByLimit(ctx context.Context, limit uint64) (uint64, error)

	CheckAccountRoleById(ctx context.Context, id string, expectedRole vo.AccountRoleType) (bool, error)
}
