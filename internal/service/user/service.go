package auth

import (
	"github.com/kenyako/auth/internal/client/db"
	"github.com/kenyako/auth/internal/repository"
	"github.com/kenyako/auth/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

func NewService(userRepository repository.UserRepository, txManager db.TxManager) service.UserService {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
	}
}
