package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	userAPI "github.com/kenyako/auth/internal/api/user"
	"github.com/kenyako/auth/internal/model"
	serviceMock "github.com/kenyako/auth/internal/service/mocks"
	desc "github.com/kenyako/auth/pkg/auth_v1"
	"github.com/stretchr/testify/require"
)

func TestImplementation_SuccessCreate(t *testing.T) {
	type mocker func() *serviceMock.UserService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()

		id = gofakeit.Int64()

		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, false, true, true, false, 8)
		role     = gofakeit.RandString([]string{"USER", "ADMIN"})

		req = &desc.CreateRequest{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: password,
			Role:            desc.UserRole(desc.UserRole_value[role]),
		}

		info = &model.UserCreate{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: password,
			Role:            role,
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)

	tests := []struct {
		name string
		args args
		want *desc.CreateResponse
		err  error
		mock mocker
	}{
		{
			name: "success user create",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			mock: func() *serviceMock.UserService {
				mockService := serviceMock.NewUserService(t)

				mockService.On("Create", ctx, info).Return(id, nil)

				return mockService
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {

			authServiceMock := tt.mock()

			api := userAPI.NewImplementation(authServiceMock)

			result, err := api.Create(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}

}

func TestImplementation_FailCreate(t *testing.T) {
	type mocker func() *serviceMock.UserService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()

		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, false, true, true, false, 8)
		role     = gofakeit.RandString([]string{"USER", "ADMIN"})

		serviceErr = fmt.Errorf("failed to create user")

		req = &desc.CreateRequest{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: password,
			Role:            desc.UserRole(desc.UserRole_value[role]),
		}

		info = &model.UserCreate{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: password,
			Role:            role,
		}
	)

	tests := []struct {
		name string
		args args
		want *desc.CreateResponse
		err  error
		mock mocker
	}{
		{
			name: "fail user create",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			mock: func() *serviceMock.UserService {
				mockService := serviceMock.NewUserService(t)

				mockService.On("Create", ctx, info).Return(int64(0), serviceErr)

				return mockService
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {

			authServiceMock := tt.mock()

			api := userAPI.NewImplementation(authServiceMock)

			result, err := api.Create(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
