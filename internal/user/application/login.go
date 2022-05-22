package application

import (
	"context"
	"errors"
	"makretplace/internal/user/domain/repository"
	"makretplace/internal/user/domain/service"
)

type LoginUseCase struct {
	repository repository.UserRepository
	service    service.UserService
}

func NewLoginUseCase(
	repository repository.UserRepository,
	service service.UserService,
) *LoginUseCase {
	return &LoginUseCase{
		repository: repository,
		service:    service,
	}
}

type LoginUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginUseCase) Execute(ctx context.Context, user LoginUserParams) (string, error) {
	item, err := l.repository.GetByEmail(ctx, user.Email)
	if err != nil {
		return "", err
	}

	if item.Password != l.service.GeneratePasswordHash(user.Password) {
		return "", errors.New("password incorrest")
	}

	return l.service.GenerateToken(item.ID)
}
