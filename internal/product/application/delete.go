package application

import (
	"context"
	"makretplace/internal/product/domain/repository"

	"github.com/google/uuid"
)

type DeleteUseCase struct {
	repository repository.ProductRepository
}

func NewDeleteUseCase(
	repository repository.ProductRepository,
) *DeleteUseCase {
	return &DeleteUseCase{
		repository: repository,
	}
}

func (d *DeleteUseCase) Execute(ctx context.Context, productId string) error {
	product, err := uuid.Parse(productId)
	if err != nil {
		return err
	}
	return d.repository.Delete(ctx, product)
}
