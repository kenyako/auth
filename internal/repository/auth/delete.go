package userrepo

import (
	"context"

	"github.com/Masterminds/squirrel"
)

func (r *repo) Delete(ctx context.Context, id int64) error {

	builderDelete := r.qb.Delete(table).
		Where(squirrel.Eq{idColumn: id})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
