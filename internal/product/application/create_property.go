package application

import (
	"context"
	"makretplace/internal/product/domain/model"
	"makretplace/internal/product/domain/repository"

	"github.com/google/uuid"
)

type CreatePropertyUseCase struct {
	repository repository.ProductPropertyRepository
}

func NewCreatePropertyUseCase(
	repository repository.ProductPropertyRepository,
) *CreatePropertyUseCase {
	return &CreatePropertyUseCase{
		repository: repository,
	}
}

func (cp *CreatePropertyUseCase) Execute(ctx context.Context, productId string, params model.CreateProductPropertyParams) (uuid.UUID, error) {
	product, err := uuid.Parse(productId)
	if err != nil {
		return uuid.Nil, err
	}
	return cp.repository.Create(ctx, product, params)
}
