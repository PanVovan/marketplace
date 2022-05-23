package application

import (
	"context"
	"makretplace/internal/brand/domain/repository"

	"github.com/google/uuid"
)

type DeleteUseCase struct {
	repository repository.BrandRepository
}

func NewDeleteUseCase(
	repository repository.BrandRepository,
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
