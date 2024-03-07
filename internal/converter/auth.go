package converter

import (
	"github.com/kenyako/auth/internal/model"
	desc "github.com/kenyako/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserCreateFromDesc(data *desc.CreateRequest) *model.UserCreate {
	return &model.UserCreate{
		Name:            data.Name,
		Email:           data.Email,
		Password:        data.Password,
		PasswordConfirm: data.PasswordConfirm,
		Role:            data.Role.String(),
	}
}

func ToUserFromService(data *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if data.UpdatedAt.Valid {
		updatedAt = timestamppb.New(data.UpdatedAt.Time)
	}

	return &desc.User{
		Id:              data.ID,
		Name:            data.Name,
		Email:           data.Email,
		Password:        data.Password,
		PasswordConfirm: data.PasswordConfirm,
		Role:            desc.UserRole(desc.UserRole_value[data.Role]),
		CreatedAt:       timestamppb.New(data.CreatedAt),
		UpdatedAt:       updatedAt,
	}
}

func ToUserUpdateFromDesc(data *desc.UpdateRequest) *model.UserUpdate {
	name := data.GetName()
	email := data.GetEmail()
	role := data.GetRole().String()

	return &model.UserUpdate{
		ID:    data.GetId(),
		Name:  &name,
		Email: &email,
		Role:  &role,
	}
}
