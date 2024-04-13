package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	userAPI "github.com/kenyako/auth/internal/api/user"
	serviceMock "github.com/kenyako/auth/internal/service/mocks"
	desc "github.com/kenyako/auth/pkg/auth_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestImplementation_SuccessDelete(t *testing.T) {

	type mocker func() *serviceMock.UserService

	type args struct {
		ctx context.Context
		req *desc.DeleteRequest
	}

	var (
		ctx = context.Background()

		id = gofakeit.Int64()

		req = &desc.DeleteRequest{
			Id: id,
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
			name: "success user delete",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			mock: func() *serviceMock.UserService {
				mockService := serviceMock.NewUserService(t)

				mockService.On("Delete", ctx, req.Id).Return(nil)

				return mockService
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			authServiceMock := tt.mock()

			api := userAPI.NewImplementation(authServiceMock)

			result, err := api.Delete(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}

func TestImplementation_FailDelete(t *testing.T) {

	type mocker func() *serviceMock.UserService

	type args struct {
		ctx context.Context
		req *desc.DeleteRequest
	}

	var (
		ctx = context.Background()

		id = gofakeit.Int64()

		serviceErr = fmt.Errorf("failed to delete user")

		req = &desc.DeleteRequest{
			Id: id,
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
			name: "fail user delete",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			mock: func() *serviceMock.UserService {
				mockService := serviceMock.NewUserService(t)

				mockService.On("Delete", ctx, req.Id).Return(serviceErr)

				return mockService
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			authServiceMock := tt.mock()

			api := userAPI.NewImplementation(authServiceMock)

			result, err := api.Delete(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
