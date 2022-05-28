package application

import (
	"context"
	"makretplace/internal/product/domain/model"
	"makretplace/internal/product/domain/repository"

	"github.com/google/uuid"
)

type GetOneUseCase struct {
	repository repository.ProductRepository
}

func NewGetOneUseCase(
	repository repository.ProductRepository,
) *GetOneUseCase {
	return &GetOneUseCase{
		repository: repository,
	}
}

func (gone *GetOneUseCase) Execute(ctx context.Context, productId string) (model.Product, error) {
	product, err := uuid.Parse(productId)
	if err != nil {
		return model.Product{}, err
	}
	return gone.repository.GetOne(ctx, product)
}
