package userrepo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kenyako/platform_common/pkg/postgres"
)

func (r *repo) Delete(ctx context.Context, id int64) error {

	builderDelete := sq.Delete(table).
		Where(sq.Eq{idColumn: id})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	q := postgres.Query{
		Name:     "auth_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args)
	if err != nil {
		return err
	}

	return nil
}
