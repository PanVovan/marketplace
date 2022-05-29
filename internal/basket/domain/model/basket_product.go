package model

import (
	"makretplace/internal/product/domain/model"

	"github.com/google/uuid"
)

type BasketProduct struct {
	ID       uuid.UUID
	Product  model.ProductInfo
	Image    model.ProductImage
	Quantity int32
}

type CreateBasketProduct struct {
	Quantity int32
}

type UpdateBasketProduct struct {
	Quantity int32
}
