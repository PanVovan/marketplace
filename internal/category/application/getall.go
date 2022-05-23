package application

import (
	"context"
	"makretplace/internal/category/domain/model"
	"makretplace/internal/category/domain/repository"
)

type GetAllUseCase struct {
	repository repository.CategoryRepository
}

func NewGetAllUseCase(
	repository repository.CategoryRepository,
) *GetAllUseCase {
	return &GetAllUseCase{
		repository: repository,
	}
}

func (ga *GetAllUseCase) Execute(ctx context.Context, limit, page int32) ([]model.Category, error) {
	return ga.repository.GetAll(ctx, limit, page)
}
