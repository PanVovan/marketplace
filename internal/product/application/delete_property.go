package application

import (
	"context"
	"makretplace/internal/product/domain/repository"

	"github.com/google/uuid"
)

type DeletePropertyUseCase struct {
	repository repository.ProductPropertyRepository
}

func NewDeletePropertyUseCase(
	repository repository.ProductPropertyRepository,
) *DeletePropertyUseCase {
	return &DeletePropertyUseCase{
		repository: repository,
	}
}

func (d *DeletePropertyUseCase) Execute(ctx context.Context, propertyId string) error {
	property, err := uuid.Parse(propertyId)
	if err != nil {
		return err
	}
	return d.repository.Delete(ctx, property)
}
