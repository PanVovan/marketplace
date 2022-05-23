package application

import (
	"context"
	"makretplace/internal/brand/domain/repository"

	"github.com/google/uuid"
)

type CreateUseCase struct {
	repository repository.BrandRepository
}

func NewCreateUseCase(
	repository repository.BrandRepository,
) *CreateUseCase {
	return &CreateUseCase{
		repository: repository,
	}
}

func (ga *CreateUseCase) Execute(ctx context.Context, name string) (uuid.UUID, error) {
	return ga.repository.Create(ctx, name)
}
