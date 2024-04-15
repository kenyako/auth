package user

import (
	"github.com/kenyako/auth/internal/repository"
	"github.com/kenyako/auth/internal/service"
	"github.com/kenyako/platform_common/pkg/postgres"
)

type serv struct {
	userRepository repository.UserRepository
	txManager      postgres.TxManager
}

func NewService(userRepository repository.UserRepository, txManager postgres.TxManager) service.UserService {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
	}
}
