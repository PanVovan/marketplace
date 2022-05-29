package repository

import (
	"context"
	"makretplace/internal/basket/domain/model"

	"github.com/google/uuid"
)

type BasketRepository interface {
	GetOne(ctx context.Context, basketId uuid.UUID) (model.Basket, error)
	Create(ctx context.Context, userId *uuid.UUID) (uuid.UUID, error)
	Append(ctx context.Context, basketId, productId uuid.UUID, quantity int32) (model.Basket, error)
	Increment(ctx context.Context, basketId, productId uuid.UUID, quantity int32) (model.Basket, error)
	Decrement(ctx context.Context, basketId, productId uuid.UUID, quantity int32) (model.Basket, error)
	Remove(ctx context.Context, basketId, productId uuid.UUID) (model.Basket, error)
	Clear(ctx context.Context, basketId uuid.UUID) (model.Basket, error)
	Delete(ctx context.Context, basketId uuid.UUID) error
}
