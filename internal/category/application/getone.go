package application

import (
	"context"
	"makretplace/internal/category/domain/model"
	"makretplace/internal/category/domain/repository"

	"github.com/google/uuid"
)

type GetOneUseCase struct {
	repository repository.CategoryRepository
}

func NewGetOneUseCase(
	repository repository.CategoryRepository,
) *GetOneUseCase {
	return &GetOneUseCase{
		repository: repository,
	}
}

func (gone *GetOneUseCase) Execute(ctx context.Context, categoryId string) (model.Category, error) {
	category, err := uuid.Parse(categoryId)
	if err != nil {
		return model.Category{}, err
	}
	return gone.repository.GetOne(ctx, category)
}
