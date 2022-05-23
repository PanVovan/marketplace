package application

import (
	"context"
	"makretplace/internal/category/domain/repository"

	"github.com/google/uuid"
)

type CreateUseCase struct {
	repository repository.CategoryRepository
}

func NewCreateUseCase(
	repository repository.CategoryRepository,
) *CreateUseCase {
	return &CreateUseCase{
		repository: repository,
	}
}

func (ga *CreateUseCase) Execute(ctx context.Context, name string) (uuid.UUID, error) {
	return ga.repository.Create(ctx, name)
}
