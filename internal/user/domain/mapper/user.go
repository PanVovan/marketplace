package mapper

import (
	"makretplace/internal/user/domain/model"
	"makretplace/internal/user/infrastructure/database/sqlc"
)

type UserMapper struct{}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (u *UserMapper) FromDTOToEntity(userDto sqlc.User) model.User {
	return model.User{
		ID:       userDto.ID,
		Email:    userDto.Email,
		Password: userDto.Password,
		Name:     userDto.Name,
	}
}

func (u *UserMapper) FromEntityToDTO(user model.User) sqlc.User {
	return sqlc.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
	}
}
