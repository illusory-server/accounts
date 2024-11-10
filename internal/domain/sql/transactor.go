package sql

import "context"

type Transactor interface {
	WithinTransaction(context.Context, func(ctx context.Context) error) error
}

type TransactionController interface {
	InjectTransaction(ctx context.Context, tx QueryExecutor) context.Context
	ExtractTransaction(ctx context.Context) QueryExecutor
}
