package auth

import (
	"context"

	"github.com/kenyako/auth/internal/converter"
	desc "github.com/kenyako/auth/pkg/auth_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		User: converter.ToUserFromService(user),
	}, nil
}
