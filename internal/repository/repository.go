package repository

import (
	"context"

	"github.com/kenyako/auth/internal/model"
)

//go:generate ../../bin/mockery --name=UserRepository --output=./mocks
type UserRepository interface {
	Create(ctx context.Context, data *model.UserCreate) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, data *model.UserUpdate) error
	Delete(ctx context.Context, id int64) error
}
