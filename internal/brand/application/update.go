package application

import (
	"context"
	"makretplace/internal/brand/domain/repository"

	"github.com/google/uuid"
)

type UpdateUseCase struct {
	repository repository.BrandRepository
}

func NewUpdateUseCase(
	repository repository.BrandRepository,
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
