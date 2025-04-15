package dependency

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/illusory-server/accounts/internal/app/factory"
	"github.com/illusory-server/accounts/internal/app/usecase/account"
	"github.com/illusory-server/accounts/internal/infra/storage/psql"
	"github.com/illusory-server/accounts/pkg/utils"
)

type timeNow struct{}

func (t timeNow) Now() time.Time {
	return time.Now()
}

type idGenerate struct{}

func (i idGenerate) GenerateID() string {
	return uuid.New().String()
}

func (f *Factory) initUseCase(ctx context.Context) error {
	cfg := f.deps.config
	pool, err := psql.ConnectPool(ctx, &cfg.Postgres)
	if err != nil {
		return err
	}
	err = pool.Ping(ctx)
	if err != nil {
		return err
	}

	log := f.deps.log

	query, err := psql.NewAccountQuery(log, pool)
	if err != nil {
		return err
	}
	command, err := psql.NewAccountCommand(log, pool, utils.NewUUIDGenerator())
	if err != nil {
		return err
	}

	timer := timeNow{}
	idGen := idGenerate{}

	accFactory := factory.NewAccountFactory(timer, idGen)

	accUseCase, err := account.NewUseCase(log, accFactory, query, command, timer)
	if err != nil {
		return err
	}

	f.deps.accountUseCase = accUseCase

	return nil
}
