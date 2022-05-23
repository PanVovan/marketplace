package repository

import (
	"context"
	"makretplace/internal/product/domain/model"

	"github.com/google/uuid"
)

type ProductPropertyRepository interface {
	Create(ctx context.Context, productID uuid.UUID, property model.CreateProductPropertyParams) (uuid.UUID, error)
	GetOne(ctx context.Context, propertyID uuid.UUID) (model.ProductProperty, error)
	GetAll(ctx context.Context, productID uuid.UUID) ([]model.ProductProperty, error)
	Update(ctx context.Context, propertyID uuid.UUID, params model.UpdateProductPropertyParams) error
	Delete(ctx context.Context, propertyID uuid.UUID) error
}
