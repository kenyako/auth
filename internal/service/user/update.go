package user

import (
	"context"

	"github.com/kenyako/auth/internal/model"
)

func (s *serv) Update(ctx context.Context, data *model.UserUpdate) error {

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTx := s.userRepository.Update(ctx, data)

		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
