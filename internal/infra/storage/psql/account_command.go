package psql

import (
	"context"
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/illusory-server/accounts/internal/domain/aggregate"
	"github.com/illusory-server/accounts/internal/domain/event"
	"github.com/illusory-server/accounts/internal/domain/repository"
	"github.com/illusory-server/accounts/pkg/errors/codex"
	"github.com/illusory-server/accounts/pkg/errors/errx"
	"github.com/illusory-server/accounts/pkg/logger"
	"github.com/illusory-server/accounts/pkg/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

var _ repository.AccountCommand = (*AccountCommand)(nil)

const (
	CreateAccountQuery = `
		INSERT INTO accounts (id, first_name, last_name, email, role, nickname, password, updated_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`
	CreateEventQuery = `
		INSERT INTO account_events (id, account_id, event_type, event_data, timestamp, aggregate_version)
		VALUES ($1, $2, $3, $4, $5, $6);
	`
	DeleteAccountByIDQuery       = "DELETE FROM accounts WHERE id = $1"
	DeleteAccountByNicknameQuery = "DELETE FROM accounts WHERE nickname = $1"
	DeleteAccountByEmailQuery    = "DELETE FROM accounts WHERE email = $1"

	UpdateAccountNicknameQuery   = "UPDATE accounts SET nickname = $2, version = version+1 WHERE id = $1 AND version = $3"
	UpdateAccountInfoQuery       = "UPDATE accounts SET first_name = $2, last_name = $3 WHERE id = $1"
	UpdateAccountEmailQuery      = "UPDATE accounts SET email = $2 WHERE id = $1"
	UpdateAccountRoleQuery       = "UPDATE accounts SET role = $2 WHERE id = $1"
	UpdateAccountPasswordQuery   = "UPDATE accounts SET password = $2 WHERE id = $1"
	UpdateAccountAvatarLinkQuery = "UPDATE accounts SET avatar_link = $2 WHERE id = $1"
)

var (
	ErrVersionNotEqual = errx.New(codex.FailedPrecondition, "update data object not equal version")
)

type AccountCommand struct {
	db    *pgxpool.Pool
	log   logger.Logger
	idGen utils.IDGenerator
}

func NewAccountCommand(log logger.Logger, pool *pgxpool.Pool, gen utils.IDGenerator) (*AccountCommand, error) {
	res := &AccountCommand{
		db:    pool,
		log:   log,
		idGen: gen,
	}

	if err := res.Validate(); err != nil {
		return nil, err
	}

	return res, nil
}

func (a *AccountCommand) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.log, validation.Required),
	)
}

func (a *AccountCommand) Create(ctx context.Context, account *aggregate.Account) (*aggregate.Account, error) {
	account.AddEvent(event.NewAccountCreate(account.Account().ID(), account.Account(), account.Account().CreatedAt()))
	_, err := a.db.Exec(ctx, CreateAccountQuery,
		account.Account().ID(),
		account.Account().Info().FirstName(),
		account.Account().Info().LastName(),
		account.Account().Info().Email(),
		account.Account().Role(),
		account.Account().Nickname(),
		account.Account().Password(),
		account.Account().UpdatedAt(),
		account.Account().CreatedAt(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "[AccountCommand] db.Exec")
	}
	return account, nil
}

func (a *AccountCommand) CreateMany(ctx context.Context, accounts []*aggregate.Account) error {
	return nil
}

func (a *AccountCommand) Update(ctx context.Context, account *aggregate.Account) error {
	for _, ev := range account.Events() {
		switch e := ev.(type) {
		case event.AccountChangeInfo:
		case event.AccountChangeNickname:
			if err := a.updateAccountNickname(ctx, e, account.Account().Version()); err == nil {
				account.VersionIncrement()
			}
		case event.AccountChangePassword:
		case event.AccountChangeEmail:
		case event.AccountChangeRole:
		case event.AccountChangeAvatarLink:
		default:
		}
	}
	return nil
}

func (a *AccountCommand) createEventWithTx(
	ctx context.Context, tx pgx.Tx,
	id string, data []byte, ev event.Event,
	aggregateVersion uint64,
) error {
	_, err := tx.Exec(ctx, CreateEventQuery,
		a.idGen.GenerateID(),
		id,
		ev.Type(),
		data,
		ev.Timestamp(),
		aggregateVersion+1,
	)
	if err != nil {
		return errors.Wrap(err, "[AccountCommand] tx.Exec")
	}
	return nil
}

func (a *AccountCommand) updateAccountTX(
	ctx context.Context,
	txAction func(ctx context.Context, tx pgx.Tx) error,
) error {
	var (
		tx  pgx.Tx
		err error
	)
	tx, err = a.db.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.RepeatableRead,
	})
	defer func() {
		if err != nil {
			e := tx.Rollback(ctx)
			if e != nil { // nolint:staticcheck
				// TODO - alert
			}
		}
	}()
	if err != nil {
		return errors.Wrap(err, "[AccountCommand] db.BeginTx")
	}
	err = txAction(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "[AccountCommand] txAction")
	}
	err = tx.Commit(ctx)
	if err != nil {
		return errors.Wrap(err, "[AccountCommand] tx.Commit")
	}
	return nil
}

func (a *AccountCommand) updateAccountNickname(
	ctx context.Context,
	nicknameEvent event.AccountChangeNickname,
	aggregateVersion uint64,
) error {
	type eventData struct {
		Nickname string `json:"nickname"`
	}
	jsonData, err := json.Marshal(eventData{Nickname: nicknameEvent.Nickname()})
	if err != nil {
		return errors.Wrap(err, "[AccountCommand] json.Marshal")
	}
	err = a.updateAccountTX(ctx, func(ctx context.Context, tx pgx.Tx) error {
		tag, err := tx.Exec(
			ctx, UpdateAccountNicknameQuery,
			nicknameEvent.ID().Value(), nicknameEvent.Nickname(),
			aggregateVersion,
		)
		if err != nil {
			return errors.Wrap(err, "[AccountCommand] tx.Exec")
		}
		if tag.RowsAffected() == 0 {
			return ErrVersionNotEqual
		}
		err = a.createEventWithTx(ctx, tx, nicknameEvent.ID().Value(), jsonData, nicknameEvent, aggregateVersion)
		if err != nil {
			return err // nolint:wrapcheck
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "[AccountCommand] updateAccountTX")
	}

	return nil
}

func (a *AccountCommand) DeleteById(ctx context.Context, id string) error {
	_, err := a.db.Exec(ctx, DeleteAccountByIDQuery, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errx.WrapWithCode(err, codex.NotFound, "[AccountCommand] db.Exec(1)")
		}
		return errx.WrapWithCode(err, codex.Internal, "[AccountCommand] db.Exec(2)")
	}
	return nil
}

func (a *AccountCommand) DeleteByEmail(ctx context.Context, email string) error {
	_, err := a.db.Exec(ctx, DeleteAccountByEmailQuery, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errx.WrapWithCode(err, codex.NotFound, "[AccountCommand] db.Exec(1)")
		}
		return errx.WrapWithCode(err, codex.Internal, "[AccountCommand] db.Exec(2)")
	}
	return nil
}

func (a *AccountCommand) DeleteByNickname(ctx context.Context, nickname string) error {
	_, err := a.db.Exec(ctx, DeleteAccountByNicknameQuery, nickname)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errx.WrapWithCode(err, codex.NotFound, "[AccountCommand] db.Exec(1)")
		}
		return errx.WrapWithCode(err, codex.Internal, "[AccountCommand] db.Exec(2)")
	}
	return nil
}
