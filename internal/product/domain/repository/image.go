package repository

import (
	"context"
	"makretplace/internal/product/domain/model"

	"github.com/google/uuid"
)

type ProductImageRepository interface {
	Create(ctx context.Context, productID uuid.UUID, filename string) (uuid.UUID, error)
	GetOne(ctx context.Context, imageID uuid.UUID) (model.ProductImage, error)
	GetAll(ctx context.Context, productID uuid.UUID) ([]model.ProductImage, error)
	Update(ctx context.Context, imageID uuid.UUID, filename string) error
	Delete(ctx context.Context, imageID uuid.UUID) error
}
