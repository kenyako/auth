package auth

import (
	"context"

	"github.com/kenyako/auth/internal/converter"
	desc "github.com/kenyako/auth/pkg/auth_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.authService.Create(ctx, converter.ToUserCreateFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
