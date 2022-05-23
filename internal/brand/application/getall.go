package application

import (
	"context"
	"makretplace/internal/brand/domain/model"
	"makretplace/internal/brand/domain/repository"
)

type GetAllUseCase struct {
	repository repository.BrandRepository
}

func NewGetAllUseCase(
	repository repository.BrandRepository,
) *GetAllUseCase {
	return &GetAllUseCase{
		repository: repository,
	}
}

func (ga *GetAllUseCase) Execute(ctx context.Context, limit, page int32) ([]model.Brand, error) {
	return ga.repository.GetAll(ctx, limit, page)
}
