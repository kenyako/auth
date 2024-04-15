package userrepo

import (
	"github.com/kenyako/auth/internal/client/postgres"
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
	db postgres.Client
}

func NewRepository(db postgres.Client) repository.UserRepository {
	return &repo{
		db: db,
	}
}
