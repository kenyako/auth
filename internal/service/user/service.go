package user

import (
	"github.com/kenyako/auth/internal/client/postgres"
	"github.com/kenyako/auth/internal/repository"
	"github.com/kenyako/auth/internal/service"
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
