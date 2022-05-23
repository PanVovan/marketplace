package application

import (
	"context"
	"makretplace/internal/category/domain/repository"

	"github.com/google/uuid"
)

type DeleteUseCase struct {
	repository repository.CategoryRepository
}

func NewDeleteUseCase(
	repository repository.CategoryRepository,
) *DeleteUseCase {
	return &DeleteUseCase{
		repository: repository,
	}
}

func (u *DeleteUseCase) Execute(ctx context.Context, brandId string) error {
	brand, err := uuid.Parse(brandId)
	if err != nil {
		return err
	}
	return u.repository.Delete(ctx, brand)
}
