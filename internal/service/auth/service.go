package auth

import (
	"github.com/kenyako/auth/internal/client/db"
	"github.com/kenyako/auth/internal/repository"
	"github.com/kenyako/auth/internal/service"
)

type serv struct {
	authRepository repository.AuthRepository
	txManager      db.TxManager
}

func NewService(authRepository repository.AuthRepository, txManager db.TxManager) service.AuthService {
	return &serv{
		authRepository: authRepository,
		txManager:      txManager,
	}
}
