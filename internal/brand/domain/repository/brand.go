package repository

import (
	"context"
	"makretplace/internal/brand/domain/model"

	"github.com/google/uuid"
)

type BrandRepository interface {
	Create(ctx context.Context, name string) (uuid.UUID, error)
	GetOne(ctx context.Context, brandID uuid.UUID) (model.Brand, error)
	GetAll(ctx context.Context, limit, page int32) ([]model.Brand, error)
	Update(ctx context.Context, brandID uuid.UUID, name string) error
	Delete(ctx context.Context, brandID uuid.UUID) error
}
