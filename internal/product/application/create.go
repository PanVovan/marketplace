package application

import (
	"context"
	"makretplace/internal/product/domain/model"
	"makretplace/internal/product/domain/repository"

	"github.com/google/uuid"
)

type CreateUseCase struct {
	repository repository.ProductRepository
}

func NewCreateUseCase(
	repository repository.ProductRepository,
) *CreateUseCase {
	return &CreateUseCase{
		repository: repository,
	}
}

func (ga *CreateUseCase) Execute(ctx context.Context, params model.CreateProductParams) (uuid.UUID, error) {
	return ga.repository.Create(ctx, params)
}
