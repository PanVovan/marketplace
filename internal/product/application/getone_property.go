package application

import (
	"context"
	"makretplace/internal/product/domain/model"
	"makretplace/internal/product/domain/repository"

	"github.com/google/uuid"
)

type GetOnePropertyUseCase struct {
	repository repository.ProductPropertyRepository
}

func NewGetOnePropertyUseCase(
	repository repository.ProductPropertyRepository,
) *GetOnePropertyUseCase {
	return &GetOnePropertyUseCase{
		repository: repository,
	}
}

func (gone *GetOnePropertyUseCase) Execute(ctx context.Context, propertyId string) (model.ProductProperty, error) {
	property, err := uuid.Parse(propertyId)
	if err != nil {
		return model.ProductProperty{}, err
	}
	return gone.repository.GetOne(ctx, property)
}
