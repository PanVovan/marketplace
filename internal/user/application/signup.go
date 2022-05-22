package application

import (
	"context"
	"makretplace/internal/user/domain/model"
	"makretplace/internal/user/domain/repository"
	"makretplace/internal/user/domain/service"

	"github.com/google/uuid"
)

type SignUpUseCase struct {
	repository repository.UserRepository
	service    service.UserService
}

func NewSignUpUseCase(
	repository repository.UserRepository,
	service service.UserService,
) *SignUpUseCase {
	return &SignUpUseCase{
		repository: repository,
		service:    service,
	}
}

func (s *SignUpUseCase) Execute(ctx context.Context, user model.CreateUserParams) (string, uuid.UUID, error) {
	user.Password = s.service.GeneratePasswordHash(user.Password)
	userId, err := s.repository.Create(ctx, user)
	if err != nil {
		return "", uuid.Nil, err
	}
	token, err := s.service.GenerateToken(userId)
	return token, userId, err
}
