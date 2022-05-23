package application

import (
	"context"
	"makretplace/internal/category/domain/repository"

	"github.com/google/uuid"
)

type UpdateUseCase struct {
	repository repository.CategoryRepository
}

func NewUpdateUseCase(
	repository repository.CategoryRepository,
) *UpdateUseCase {
	return &UpdateUseCase{
		repository: repository,
	}
}

func (u *UpdateUseCase) Execute(ctx context.Context, brandId string, name string) error {
	brand, err := uuid.Parse(brandId)
	if err != nil {
		return err
	}
	return u.repository.Update(ctx, brand, name)
}
