package dependency

import (
	"context"

	"github.com/illusory-server/accounts/internal/app/usecase/account"
	"github.com/illusory-server/accounts/internal/infra/config"
	"github.com/illusory-server/accounts/pkg/logger/log"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Dependency struct {
	config *config.Config
	log    *log.Log

	accountUseCase *account.UseCase
}

func (d Dependency) AccountUseCase() *account.UseCase {
	return d.accountUseCase
}

func (d Dependency) Config() *config.Config {
	return d.config
}

func (d Dependency) Logger() *log.Log {
	return d.log
}

type Factory struct {
	deps     Dependency
	psqlPool *pgxpool.Pool
}

func NewFactory() *Factory {
	ctx := context.Background()
	f := &Factory{}
	m := []func(context.Context) error{
		f.initConfigAndLogger,
		f.initUseCase,
	}
	for _, fn := range m {
		if err := fn(ctx); err != nil {
			panic(err)
		}
	}
	return f
}

func (f *Factory) Dependency() Dependency {
	return f.deps
}

func (f *Factory) Close() {
	f.psqlPool.Close()
}
