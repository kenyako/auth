package converter

import (
	"github.com/kenyako/auth/internal/model"
	coremodel "github.com/kenyako/auth/internal/repository/user/model"
)

func ToUserFromRepo(user *coremodel.User) *model.User {

	return &model.User{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		Password:        user.Password,
		PasswordConfirm: user.PasswordConfirm,
		Role:            user.Role,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
	}
}
