package tests

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"testing"

// 	"github.com/brianvoe/gofakeit"
// 	"github.com/gojuno/minimock/v3"
// 	"github.com/kenyako/auth/internal/model"
// 	"github.com/kenyako/auth/internal/repository"
// 	repoMocks "github.com/kenyako/auth/internal/repository/mocks"
// 	"github.com/kenyako/auth/internal/service/auth"
// 	"github.com/stretchr/testify/require"
// )

// func TestGet(t *testing.T) {

// 	type authRepositoryMockFunc func(mc *minimock.Controller) repository.AuthRepository

// 	type args struct {
// 		ctx   context.Context
// 		reqId int64
// 	}

// 	var (
// 		ctx = context.Background()
// 		mc  = minimock.NewController(t)

// 		reqId = gofakeit.Int64()

// 		name       = gofakeit.Name()
// 		email      = gofakeit.Email()
// 		password   = gofakeit.Password(true, false, true, true, false, 9)
// 		role       = gofakeit.RandString([]string{"USER", "ADMIN"})
// 		created_at = gofakeit.Date()
// 		updated_at = gofakeit.Date()

// 		res = &model.User{
// 			ID:              reqId,
// 			Name:            name,
// 			Email:           email,
// 			Password:        password,
// 			PasswordConfirm: password,
// 			Role:            role,
// 			CreatedAt:       created_at,
// 			UpdatedAt: sql.NullTime{
// 				Time:  updated_at,
// 				Valid: true,
// 			},
// 		}

// 		repoErr = fmt.Errorf("repo error")
// 	)

// 	tests := []struct {
// 		name               string
// 		args               args
// 		want               *model.User
// 		err                error
// 		authRepositoryMock authRepositoryMockFunc
// 	}{
// 		{
// 			name: "success case",
// 			args: args{
// 				ctx:   ctx,
// 				reqId: reqId,
// 			},
// 			want: res,
// 			err:  nil,
// 			authRepositoryMock: func(mc *minimock.Controller) repository.AuthRepository {
// 				mock := repoMocks.NewAuthRepositoryMock(mc)
// 				mock.GetMock.Expect(ctx, reqId).Return(res, nil)
// 				return mock
// 			},
// 		},
// 		{
// 			name: "repo error case",
// 			args: args{
// 				ctx:   ctx,
// 				reqId: reqId,
// 			},
// 			want: nil,
// 			err:  repoErr,
// 			authRepositoryMock: func(mc *minimock.Controller) repository.AuthRepository {
// 				mock := repoMocks.NewAuthRepositoryMock(mc)
// 				mock.GetMock.Expect(ctx, reqId).Return(nil, repoErr)
// 				return mock
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		tt := tt

// 		t.Run(tt.name, func(t *testing.T) {
// 			authRepositoryMock := tt.authRepositoryMock(mc)
// 			service := auth.NewMockService(authRepositoryMock)

// 			result, err := service.Get(tt.args.ctx, tt.args.reqId)

// 			require.Equal(t, tt.err, err)
// 			require.Equal(t, tt.want, result)
// 		})
// 	}
// }
