package application

import (
	"context"
	"makretplace/internal/product/domain/model"
	"makretplace/internal/product/domain/repository"

	"github.com/google/uuid"
)

type UpdateUseCase struct {
	repository repository.ProductRepository
}

func NewUpdateUseCase(
	repository repository.ProductRepository,
) *UpdateUseCase {
	return &UpdateUseCase{
		repository: repository,
	}
}

func (u *UpdateUseCase) Execute(ctx context.Context, productId string, params model.UpdateProductParams) error {

	product, err := uuid.Parse(productId)
	if err != nil {
		return err
	}
	return u.repository.Update(ctx, product, params)
}
