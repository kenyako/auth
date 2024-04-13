package userrepo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kenyako/auth/internal/client/db"
	"github.com/kenyako/auth/internal/model"
	"github.com/kenyako/auth/internal/repository/user/converter"
	coremodel "github.com/kenyako/auth/internal/repository/user/model"
)

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {

	builderSelect := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(table).
		Where(sq.Eq{
			idColumn: id,
		}).
		Limit(1)

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "auth_repository.Get",
		QueryRaw: query,
	}

	var user coremodel.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}
