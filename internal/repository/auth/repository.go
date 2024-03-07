package userrepo

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kenyako/auth/internal/repository"
)

const (
	table = "users"
)

const (
	idColumn              = "id"
	nameColumn            = "name"
	emailColumn           = "email"
	passwordColumn        = "password"
	passwordConfirmColumn = "password_confirm"
	roleColumn            = "role"
	createdAtColumn       = "created_at"
	updatedAtColumn       = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
	qb squirrel.StatementBuilderType
}

func NewRepository(db *pgxpool.Pool) repository.AuthRepository {
	return &repo{
		db: db,
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
