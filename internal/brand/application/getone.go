package application

import (
	"context"
	"makretplace/internal/brand/domain/model"
	"makretplace/internal/brand/domain/repository"

	"github.com/google/uuid"
)

type GetOneUseCase struct {
	repository repository.BrandRepository
}

func NewGetOneUseCase(
	repository repository.BrandRepository,
) *GetOneUseCase {
	return &GetOneUseCase{
		repository: repository,
	}
}

func (gone *GetOneUseCase) Execute(ctx context.Context, brandId string) (model.Brand, error) {
	brand, err := uuid.Parse(brandId)
	if err != nil {
		return model.Brand{}, err
	}
	return gone.repository.GetOne(ctx, brand)
}
