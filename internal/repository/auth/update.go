package userrepo

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/kenyako/auth/internal/client/db"
	"github.com/kenyako/auth/internal/model"
)

func (r *repo) Update(ctx context.Context, data *model.UserUpdate) error {

	builderUpdate := sq.Update(table).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: data.ID})

	if data.Name != nil {
		builderUpdate = builderUpdate.Set(nameColumn, data.Name)
	}

	if data.Email != nil {
		builderUpdate = builderUpdate.Set(emailColumn, data.Email)
	}

	if data.Role != nil {
		builderUpdate = builderUpdate.Set(roleColumn, data.Role)
	}

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "auth_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
