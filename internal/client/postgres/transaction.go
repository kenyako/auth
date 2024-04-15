package postgres

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

//go:generate ../../../bin/mockery --name=Tx --output=./mocks
type Tx interface {
	Begin(ctx context.Context) (pgx.Tx, error)

	Commit(ctx context.Context) error

	Rollback(ctx context.Context) error

	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	LargeObjects() pgx.LargeObjects

	Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error)

	Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row

	Conn() *pgx.Conn
}

type txKey struct{}

//go:generate ../../../bin/mockery --name=TxManager --output=./mocks
type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

type Handler func(ctx context.Context) error

type txManager struct {
	db Postgres
}

func NewTransactionManager(db Postgres) TxManager {
	return &txManager{
		db: db,
	}
}

func (m *txManager) transaction(ctx context.Context, opts pgx.TxOptions, fn Handler) (err error) {

	tx, ok := ExtractTx(ctx)
	if ok {
		return fn(ctx)
	}

	tx, err = m.db.BeginTx(ctx, opts)
	if err != nil {
		return errors.Wrap(err, "can't begin transaction")
	}

	ctx = InjectTx(ctx, tx)

	defer func() {

		if r := recover(); r != nil {
			err = errors.Errorf("panic recovered: #{r}")
		}

		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Wrapf(err, "errRollback: #{errRollback}")
			}

			return
		}

		if err == nil {
			err = tx.Commit(ctx)
			if err != nil {
				err = errors.Wrap(err, "tx commit failed")
			}
		}
	}()

	if err = fn(ctx); err != nil {
		err = errors.Wrap(err, "failed executing code inside transaction")
	}

	return err
}

func (m *txManager) ReadCommitted(ctx context.Context, f Handler) error {
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	return m.transaction(ctx, txOpts, f)
}

func InjectTx(ctx context.Context, tx Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func ExtractTx(ctx context.Context) (Tx, bool) {
	tx, ok := ctx.Value(txKey{}).(Tx)

	return tx, ok
}
