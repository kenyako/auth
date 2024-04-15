package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v4"
	"github.com/kenyako/auth/internal/model"
	"github.com/kenyako/auth/internal/repository"
	"github.com/kenyako/auth/internal/service/user"
	"github.com/kenyako/platform_common/pkg/postgres"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	usermocks "github.com/kenyako/auth/internal/repository/mocks"
	postgresmocks "github.com/kenyako/platform_common/pkg/postgres/mocks"
)

func TestService_SuccessGet(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		req int64
	}

	type mocker struct {
		userRepo  repository.UserRepository
		txManager postgres.TxManager
	}

	var (
		ctx = context.Background()

		txOpts = pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}

		id = gofakeit.Int64()

		name       = gofakeit.Name()
		email      = gofakeit.Email()
		password   = gofakeit.Password(true, false, true, true, false, 9)
		role       = gofakeit.RandString([]string{"USER", "ADMIN"})
		created_at = gofakeit.Date()
		updated_at = gofakeit.Date()

		res = &model.User{
			ID:              id,
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: password,
			Role:            role,
			CreatedAt:       created_at,
			UpdatedAt: sql.NullTime{
				Time:  updated_at,
				Valid: true,
			},
		}
	)

	tests := []struct {
		name string
		args args
		want *model.User
		err  error
		mock func(tt args) mocker
	}{
		{
			name: "success repo user get",
			args: args{
				ctx: ctx,
				req: id,
			},
			want: res,
			err:  nil,
			mock: func(tt args) mocker {

				tx := postgresmocks.NewTx(t)

				txCtx := postgres.InjectTx(tt.ctx, tx)

				tx.On("Commit", txCtx).Return(nil)

				db := postgresmocks.NewPostgres(t)
				db.On("BeginTx", tt.ctx, txOpts).Return(tx, nil)

				txManager := postgres.NewTransactionManager(db)

				userRepo := usermocks.NewUserRepository(t)
				userRepo.On("Get", txCtx, tt.req).Return(res, nil)

				return mocker{
					userRepo:  userRepo,
					txManager: txManager,
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockerArgs := tt.mock(tt.args)

			service := user.NewService(mockerArgs.userRepo, mockerArgs.txManager)

			result, err := service.Get(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}

func TestService_FailGet(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		req int64
	}

	type mocker struct {
		userRepo  repository.UserRepository
		txManager postgres.TxManager
	}

	var (
		ctx = context.Background()

		txOpts = pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}

		repoErr = fmt.Errorf("failed to user get")

		id = gofakeit.Int64()
	)

	tests := []struct {
		name string
		args args
		want *model.User
		err  error
		mock func(tt args) mocker
	}{
		{
			name: "fail repo user get",
			args: args{
				ctx: ctx,
				req: id,
			},
			want: nil,
			err:  repoErr,
			mock: func(tt args) mocker {

				tx := postgresmocks.NewTx(t)

				txCtx := postgres.InjectTx(tt.ctx, tx)

				tx.On("Rollback", txCtx).Return(nil)

				db := postgresmocks.NewPostgres(t)
				db.On("BeginTx", tt.ctx, txOpts).Return(tx, nil)

				txManager := postgres.NewTransactionManager(db)

				userRepo := usermocks.NewUserRepository(t)
				userRepo.On("Get", txCtx, tt.req).Return(nil, repoErr)

				return mocker{
					userRepo:  userRepo,
					txManager: txManager,
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockerArgs := tt.mock(tt.args)

			service := user.NewService(mockerArgs.userRepo, mockerArgs.txManager)

			result, err := service.Get(tt.args.ctx, tt.args.req)

			require.Error(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}

func TestService_FailProcessTxUserGet(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		req int64
	}

	type mocker struct {
		userRepo  repository.UserRepository
		txManager postgres.TxManager
	}

	var (
		ctx = context.Background()

		repoErr = fmt.Errorf("failed to transaction start")

		id = gofakeit.Int64()
	)

	tests := []struct {
		name string
		args args
		want *model.User
		err  error
		mock func(tt args) mocker
	}{
		{
			name: "fail to start tx (Read Committed returned error)",
			args: args{
				ctx: ctx,
				req: id,
			},
			want: nil,
			err:  repoErr,
			mock: func(tt args) mocker {

				txManager := postgresmocks.NewTxManager(t)

				txManager.On("ReadCommitted", tt.ctx, mock.AnythingOfType("postgres.Handler")).Return(repoErr)

				return mocker{
					userRepo:  nil,
					txManager: txManager,
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockerArgs := tt.mock(tt.args)

			service := user.NewService(mockerArgs.userRepo, mockerArgs.txManager)

			result, err := service.Get(tt.args.ctx, tt.args.req)

			require.Error(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
