package auth

import (
	"github.com/kenyako/auth/internal/service"
	desc "github.com/kenyako/auth/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedUserAPIServer
	authService service.AuthService
}

func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
