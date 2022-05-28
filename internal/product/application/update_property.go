package application

import (
	"context"
	"makretplace/internal/product/domain/model"
	"makretplace/internal/product/domain/repository"

	"github.com/google/uuid"
)

type UpdatePropertyUseCase struct {
	repository repository.ProductPropertyRepository
}

func NewUpdatePropertyUseCase(
	repository repository.ProductPropertyRepository,
) *UpdatePropertyUseCase {
	return &UpdatePropertyUseCase{
		repository: repository,
	}
}

func (u *UpdatePropertyUseCase) Execute(ctx context.Context, propertyId string, params model.UpdateProductPropertyParams) error {

	property, err := uuid.Parse(propertyId)
	if err != nil {
		return err
	}
	return u.repository.Update(ctx, property, params)
}
