package repository

import (
	"context"
	"makretplace/internal/category/domain/model"

	"github.com/google/uuid"
)

type CategoryRepository interface {
	Create(ctx context.Context, name string) (uuid.UUID, error)
	GetOne(ctx context.Context, categoryID uuid.UUID) (model.Category, error)
	GetAll(ctx context.Context, limit, page int32) ([]model.Category, error)
	Update(ctx context.Context, categoryID uuid.UUID, name string) error
	Delete(ctx context.Context, categoryID uuid.UUID) error
}
