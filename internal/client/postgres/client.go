package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Client interface {
	DB() Postgres
	Close() error
}

type pgClient struct {
	masterDBC Postgres
}

func NewClient(ctx context.Context, dsn string) (Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &pgClient{
		masterDBC: &pg{dbc: dbc},
	}, nil
}

func (p *pgClient) DB() Postgres {
	return p.masterDBC
}

func (p *pgClient) Close() error {
	if p.masterDBC != nil {
		p.masterDBC.Close()
	}

	return nil
}
