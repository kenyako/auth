package auth

import (
	"github.com/kenyako/auth/internal/repository"
	"github.com/kenyako/auth/internal/service"
)

type serv struct {
	authRepository repository.AuthRepository
}

func NewService(authRepository repository.AuthRepository) service.AuthService {
	return &serv{
		authRepository: authRepository,
	}
}
