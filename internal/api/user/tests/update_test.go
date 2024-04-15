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
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestImplementation_SuccessUpdate(t *testing.T) {

	type mocker func() *serviceMock.UserService

	type args struct {
		ctx context.Context
		req *desc.UpdateRequest
	}

	var (
		ctx = context.Background()

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		role_desc = desc.UserRole(desc.UserRole_value[gofakeit.RandString([]string{"USER", "ADMIN"})])
		role_str  = gofakeit.RandString([]string{"USER", "ADMIN"})

		req = &desc.UpdateRequest{
			Id:    id,
			Name:  &name,
			Email: &email,
			Role:  &role_desc,
		}

		info = &model.UserUpdate{
			ID:    id,
			Name:  &name,
			Email: &email,
			Role:  &role_str,
		}

		res = &emptypb.Empty{}
	)

	tests := []struct {
		name string
		args args
		want *emptypb.Empty
		err  error
		mock mocker
	}{
		{
			name: "success user update",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			mock: func() *serviceMock.UserService {
				mockService := serviceMock.NewUserService(t)

				mockService.On("Update", ctx, info).Return(nil)

				return mockService
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {

			authServiceMock := tt.mock()

			api := userAPI.NewImplementation(authServiceMock)

			result, err := api.Update(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}

func TestImplementation_FailUpdate(t *testing.T) {

	type mocker func() *serviceMock.UserService

	type args struct {
		ctx context.Context
		req *desc.UpdateRequest
	}

	var (
		ctx = context.Background()

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		role_desc = desc.UserRole(desc.UserRole_value[gofakeit.RandString([]string{"USER", "ADMIN"})])
		role_str  = gofakeit.RandString([]string{"USER", "ADMIN"})

		serviceErr = fmt.Errorf("failed to update user")

		req = &desc.UpdateRequest{
			Id:    id,
			Name:  &name,
			Email: &email,
			Role:  &role_desc,
		}

		info = &model.UserUpdate{
			ID:    id,
			Name:  &name,
			Email: &email,
			Role:  &role_str,
		}
	)

	tests := []struct {
		name string
		args args
		want *emptypb.Empty
		err  error
		mock mocker
	}{
		{
			name: "fail user update",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			mock: func() *serviceMock.UserService {
				mockService := serviceMock.NewUserService(t)

				mockService.On("Update", ctx, info).Return(serviceErr)

				return mockService
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {

			authServiceMock := tt.mock()

			api := userAPI.NewImplementation(authServiceMock)

			result, err := api.Update(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
