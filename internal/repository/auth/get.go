package userrepo

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/kenyako/auth/internal/model"
	"github.com/kenyako/auth/internal/repository/auth/converter"
	coremodel "github.com/kenyako/auth/internal/repository/auth/model"
)

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {

	builderSelect := r.qb.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(table).
		Where(squirrel.Eq{
			idColumn: id,
		})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	user, err := pgx.CollectOneRow(row, pgx.RowToAddrOfStructByNameLax[coremodel.User])
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(user), nil
}
