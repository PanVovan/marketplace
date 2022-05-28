package application

import (
	"context"
	"makretplace/internal/product/domain/model"
	"makretplace/internal/product/domain/repository"

	"github.com/google/uuid"
)

type GetAllPropertyUseCase struct {
	repository repository.ProductPropertyRepository
}

func NewGetAllPropertyUseCase(
	repository repository.ProductPropertyRepository,
) *GetAllPropertyUseCase {
	return &GetAllPropertyUseCase{
		repository: repository,
	}
}

func (ga *GetAllPropertyUseCase) Execute(ctx context.Context, productId string) ([]model.ProductProperty, error) {

	product, err := uuid.Parse(productId)
	if err != nil {
		return nil, err
	}
	return ga.repository.GetAll(ctx, product)
}
