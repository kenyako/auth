package userrepo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kenyako/auth/internal/model"
)

func (r *repo) Create(ctx context.Context, data *model.UserCreate) (int64, error) {
	builderInsert := r.qb.Insert(table).
		Columns(nameColumn, emailColumn, passwordColumn, passwordConfirmColumn, roleColumn).
		Values(data.Name, data.Email, data.Password, data.PasswordConfirm, data.Role).
		Suffix(fmt.Sprintf("RETURNING %s", idColumn))

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return 0, nil
	}

	id, err := pgx.CollectOneRow(rows, pgx.RowTo[int64])
	if err != nil {
		return 0, err
	}

	return id, nil
}
