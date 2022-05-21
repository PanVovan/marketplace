package model

import (
	brand_model "makretplace/internal/brand/domain/model"
	category_model "makretplace/internal/category/domain/model"
	seller_model "makretplace/internal/seller/domain/model"

	"github.com/google/uuid"
)

type ProductInfo struct {
	ID         uuid.UUID
	Name       string
	Price      float64
	Rating     float64
	Seller     seller_model.SellerInfo
	Brand      brand_model.Brand
	Categories []category_model.Category
	Amount     int32
}
