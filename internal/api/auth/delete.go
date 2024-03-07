package auth

import (
	"context"

	desc "github.com/kenyako/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.authService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
