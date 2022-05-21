package model

import (
	brand_model "makretplace/internal/brand/domain/model"
	categoty_model "makretplace/internal/category/domain/model"
	seller_model "makretplace/internal/seller/domain/model"

	"github.com/google/uuid"
)

type Product struct {
	ID         uuid.UUID
	Name       string
	Price      float64
	Rating     float64
	Brand      brand_model.Brand
	Seller     seller_model.SellerInfo
	Amount     int32
	Images     []ProductImage
	Properties []ProductProperty
	Categories []categoty_model.Category
}
