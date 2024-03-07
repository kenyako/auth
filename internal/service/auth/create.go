package auth

import (
	"context"

	"github.com/kenyako/auth/internal/model"
)

func (s *serv) Create(ctx context.Context, data *model.UserCreate) (int64, error) {

	id, err := s.authRepository.Create(ctx, data)
	if err != nil {
		return 0, err
	}

	return id, nil
}
