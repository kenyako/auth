package auth

import (
	"context"

	"github.com/kenyako/auth/internal/model"
)

func (s *serv) Update(ctx context.Context, data *model.UserUpdate) error {
	err := s.authRepository.Update(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
