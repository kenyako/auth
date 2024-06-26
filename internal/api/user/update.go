package auth

import (
	"context"

	"github.com/kenyako/auth/internal/converter"
	desc "github.com/kenyako/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := i.userService.Update(ctx, converter.ToUserUpdateFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
