package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	userAPI "github.com/kenyako/auth/internal/api/user"
	"github.com/kenyako/auth/internal/model"
	serviceMock "github.com/kenyako/auth/internal/service/mocks"
	desc "github.com/kenyako/auth/pkg/auth_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestImplementation_SuccessGet(t *testing.T) {

	type mocker func() *serviceMock.UserService

	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}

	var (
		ctx = context.Background()

		id         = gofakeit.Int64()
		name       = gofakeit.Name()
		email      = gofakeit.Email()
		password   = gofakeit.Password(true, false, true, true, false, 9)
		role       = gofakeit.RandString([]string{"USER", "ADMIN"})
		created_at = gofakeit.Date()
		updeted_at = gofakeit.Date()

		req = &desc.GetRequest{
			Id: id,
		}

		serviceRes = &model.User{
			ID:              id,
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: password,
			Role:            role,
			CreatedAt:       created_at,
			UpdatedAt: sql.NullTime{
				Time:  updeted_at,
				Valid: true,
			},
		}

		res = &desc.GetResponse{
			User: &desc.User{
				Id:              id,
				Name:            name,
				Email:           email,
				Password:        password,
				PasswordConfirm: password,
				Role:            desc.UserRole(desc.UserRole_value[role]),
				CreatedAt:       timestamppb.New(created_at),
				UpdatedAt:       timestamppb.New(updeted_at),
			},
		}
	)

	tests := []struct {
		name string
		args args
		want *desc.GetResponse
		err  error
		mock mocker
	}{
		{
			name: "success user get",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			mock: func() *serviceMock.UserService {
				mockService := serviceMock.NewUserService(t)

				mockService.On("Get", ctx, req.Id).Return(serviceRes, nil)

				return mockService
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {

			authServiceMock := tt.mock()

			api := userAPI.NewImplementation(authServiceMock)

			result, err := api.Get(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}

func TestImplementation_FailGet(t *testing.T) {

	type mocker func() *serviceMock.UserService

	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}

	var (
		ctx = context.Background()

		id = gofakeit.Int64()

		serviceErr = fmt.Errorf("failed to get user")

		req = &desc.GetRequest{
			Id: id,
		}
	)

	tests := []struct {
		name string
		args args
		want *desc.GetResponse
		err  error
		mock mocker
	}{
		{
			name: "fail user get",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			mock: func() *serviceMock.UserService {
				mockService := serviceMock.NewUserService(t)

				mockService.On("Get", ctx, req.Id).Return(nil, serviceErr)

				return mockService
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {

			authServiceMock := tt.mock()

			api := userAPI.NewImplementation(authServiceMock)

			result, err := api.Get(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
