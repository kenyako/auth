package auth

import (
	"context"

	"github.com/kenyako/auth/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {

	user, err := s.authRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
