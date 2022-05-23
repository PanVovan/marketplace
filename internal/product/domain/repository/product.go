package repository

import (
	"context"
	"makretplace/internal/product/domain/model"

	"github.com/google/uuid"
)

type GetProductsQuerySpecs struct {
	SellerID     *uuid.UUID
	CategoriesID []*uuid.UUID
	BrandID      *uuid.UUID
}
type ProductRepository interface {
	Create(ctx context.Context, property model.CreateProductParams) (uuid.UUID, error)
	GetOne(ctx context.Context, productID uuid.UUID) (model.Product, error)
	GetAll(ctx context.Context, specs GetProductsQuerySpecs, limit, page int32) ([]model.ProductInfo, error)
	Update(ctx context.Context, productID uuid.UUID, params model.UpdateProductParams) error
	Delete(ctx context.Context, productID uuid.UUID) error
}
