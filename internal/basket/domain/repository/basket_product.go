package repository

import (
	"context"
	"makretplace/internal/basket/domain/model"

	"github.com/google/uuid"
)

type BasketProductRepository interface {
	GetOne(ctx context.Context, basketId, productId uuid.UUID) (model.BasketProduct, error)
	GetAll(ctx context.Context, basketId uuid.UUID) ([]model.BasketProduct, error)
	Create(ctx context.Context, basketId uuid.UUID, create model.CreateBasketProduct) (uuid.UUID, error)
	Update(ctx context.Context, basketId uuid.UUID, update model.UpdateBasketProduct) error
	Delete(ctx context.Context, productId uuid.UUID) error
}
