package userrepo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/kenyako/auth/internal/client/db"
	"github.com/kenyako/auth/internal/model"
)

func (r *repo) Create(ctx context.Context, data *model.UserCreate) (int64, error) {
	builderInsert := sq.Insert(table).
		Columns(nameColumn, emailColumn, passwordColumn, passwordConfirmColumn, roleColumn).
		Values(data.Name, data.Email, data.Password, data.PasswordConfirm, data.Role).
		Suffix(fmt.Sprintf("RETURNING %s", idColumn))

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "auth_repository.Create",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, nil
	}

	return id, nil
}
