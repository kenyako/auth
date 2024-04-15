package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v4"
	"github.com/kenyako/auth/internal/client/postgres"
	"github.com/kenyako/auth/internal/repository"
	"github.com/kenyako/auth/internal/service/user"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	postgresmocks "github.com/kenyako/auth/internal/client/postgres/mocks"
	usermocks "github.com/kenyako/auth/internal/repository/mocks"
)

func TestService_SuccessDelete(t *testing.T) {
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
	)

	tests := []struct {
		name string
		args args
		err  error
		mock func(tt args) mocker
	}{
		{
			name: "success repo user delete",
			args: args{
				ctx: ctx,
				req: id,
			},
			err: nil,
			mock: func(tt args) mocker {

				tx := postgresmocks.NewTx(t)

				txCtx := postgres.InjectTx(tt.ctx, tx)

				tx.On("Commit", txCtx).Return(nil)

				db := postgresmocks.NewPostgres(t)
				db.On("BeginTx", tt.ctx, txOpts).Return(tx, nil)

				txManager := postgres.NewTransactionManager(db)

				userRepo := usermocks.NewUserRepository(t)
				userRepo.On("Delete", txCtx, tt.req).Return(nil)

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

			err := service.Delete(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
		})
	}
}

func TestService_FailDelete(t *testing.T) {
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

		repoErr = fmt.Errorf("failed to user delete")

		id = gofakeit.Int64()
	)

	tests := []struct {
		name string
		args args
		err  error
		mock func(tt args) mocker
	}{
		{
			name: "fail repo user delete",
			args: args{
				ctx: ctx,
				req: id,
			},
			err: repoErr,
			mock: func(tt args) mocker {

				tx := postgresmocks.NewTx(t)

				txCtx := postgres.InjectTx(ctx, tx)

				tx.On("Rollback", txCtx).Return(nil)

				db := postgresmocks.NewPostgres(t)
				db.On("BeginTx", tt.ctx, txOpts).Return(tx, nil)

				txManager := postgres.NewTransactionManager(db)

				userRepo := usermocks.NewUserRepository(t)
				userRepo.On("Delete", txCtx, tt.req).Return(repoErr)

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

			err := service.Delete(tt.args.ctx, tt.args.req)

			require.Error(t, tt.err, err)
		})
	}
}

func TestService_FailProcessTxUserDelete(t *testing.T) {
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
		err  error
		mock func(tt args) mocker
	}{
		{
			name: "fail to start tx (Read Committed returned error)",
			args: args{
				ctx: ctx,
				req: id,
			},
			err: repoErr,
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

			err := service.Delete(tt.args.ctx, tt.args.req)

			require.Error(t, tt.err, err)
		})
	}
}
