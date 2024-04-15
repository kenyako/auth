package auth

import (
	"github.com/kenyako/auth/internal/service"
	desc "github.com/kenyako/auth/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedUserAPIServer
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
