package tests

// import (
// 	"context"
// 	"fmt"
// 	"testing"

// 	"github.com/brianvoe/gofakeit"
// 	"github.com/gojuno/minimock/v3"
// 	"github.com/jackc/pgx/v4"
// 	"github.com/kenyako/auth/internal/client/db"
// 	"github.com/kenyako/auth/internal/model"
// 	"github.com/kenyako/auth/internal/repository"
// 	"github.com/kenyako/auth/internal/service/auth"
// 	"github.com/stretchr/testify/require"

// 	dbMocks "github.com/kenyako/auth/internal/client/db/mocks"
// 	repoMocks "github.com/kenyako/auth/internal/repository/mocks"
// )

// func TestCreate(t *testing.T) {

// 	type authRepositoryMockFunc func(mc *minimock.Controller) repository.AuthRepository
// 	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

// 	type args struct {
// 		ctx context.Context
// 		req *model.UserCreate
// 	}

// 	var (
// 		ctx    = context.Background()
// 		mc     = minimock.NewController(t)
// 		txOpts = pgx.TxOptions{
// 			IsoLevel: pgx.ReadCommitted,
// 		}

// 		id       = gofakeit.Int64()
// 		name     = gofakeit.Name()
// 		email    = gofakeit.Email()
// 		password = gofakeit.Password(true, false, true, true, false, 9)
// 		role     = gofakeit.RandString([]string{"USER", "ADMIN"})

// 		repoErr = fmt.Errorf("repo error")

// 		req = &model.UserCreate{
// 			Name:            name,
// 			Email:           email,
// 			Password:        password,
// 			PasswordConfirm: password,
// 			Role:            role,
// 		}
// 	)

// 	tests := []struct {
// 		name               string
// 		args               args
// 		want               int64
// 		err                error
// 		authRepositoryMock authRepositoryMockFunc
// 		txManagerMock      txManagerMockFunc
// 	}{
// 		{
// 			name: "success case",
// 			args: args{
// 				ctx: ctx,
// 				req: req,
// 			},
// 			want: id,
// 			err:  nil,
// 			authRepositoryMock: func(mc *minimock.Controller) repository.AuthRepository {
// 				mock := repoMocks.NewAuthRepositoryMock(mc)
// 				mock.CreateMock.Expect(ctx, req).Return(id, nil)
// 				return mock
// 			},
// 			txManagerMock: func(mc *minimock.Controller) db.TxManager {
// 				mockClient := dbMocks.NewClientMock(mc)

// 			},
// 		},
// 		{
// 			name: "repo error case",
// 			args: args{
// 				ctx: ctx,
// 				req: req,
// 			},
// 			want: 0,
// 			err:  repoErr,
// 			authRepositoryMock: func(mc *minimock.Controller) repository.AuthRepository {
// 				mock := repoMocks.NewAuthRepositoryMock(mc)
// 				mock.CreateMock.Expect(ctx, req).Return(0, repoErr)
// 				return mock
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		tt := tt

// 		t.Run(tt.name, func(t *testing.T) {

// 			authRepositoryMock := tt.authRepositoryMock(mc)
// 			txManagerMock := tt.txManagerMock(mc)
// 			service := auth.NewService(authRepositoryMock, txManagerMock)

// 			result, err := service.Create(tt.args.ctx, tt.args.req)

// 			require.Equal(t, tt.err, err)
// 			require.Equal(t, tt.want, result)
// 		})
// 	}
// }
