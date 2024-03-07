package userrepo

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/kenyako/auth/internal/model"
)

func (r *repo) Update(ctx context.Context, data *model.UserUpdate) error {

	builderUpdate := r.qb.Update(table).
		Set(updatedAtColumn, time.Now()).
		Where(squirrel.Eq{idColumn: data.ID})

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

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
