package auth

import (
	"context"

	"github.com/kenyako/auth/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {

	var user *model.User

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		u, errTx := s.userRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		user = u

		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}
