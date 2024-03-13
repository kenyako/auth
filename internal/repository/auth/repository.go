package userrepo

import (
	"github.com/kenyako/auth/internal/client/db"
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
	db db.Client
}

func NewRepository(db db.Client) repository.AuthRepository {
	return &repo{
		db: db,
	}
}
