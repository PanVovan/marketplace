package model

import (
	brand_model "makretplace/internal/brand/domain/model"
	categoty_model "makretplace/internal/category/domain/model"
	seller_model "makretplace/internal/seller/domain/model"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID                 `json:"id"`
	Name        string                    `json:"name"`
	Price       float64                   `json:"price"`
	Rating      float64                   `json:"rating"`
	Brand       brand_model.Brand         `json:"brand"`
	Description string                    `json:"description"`
	Seller      seller_model.SellerInfo   `json:"seller"`
	Amount      int32                     `json:"amount"`
	Images      []ProductImage            `json:"images"`
	Properties  []ProductProperty         `json:"properties"`
	Categories  []categoty_model.Category `json:"categories"`
}

type CreateProductParams struct {
	Name        string            `json:"name"`
	Price       float64           `json:"price"`
	Rating      float64           `json:"rating"`
	BrandID     *uuid.UUID        `json:"brand_id"`
	Description string            `json:"description"`
	SellerID    uuid.UUID         `json:"seller_id"`
	Amount      int32             `json:"amount"`
	Images      []ProductImage    `json:"images"`
	Properties  []ProductProperty `json:"properties"`
	Categories  []uuid.UUID       `json:"categories"`
}

type UpdateProductParams struct {
	Name        *string           `json:"name"`
	Price       *float64          `json:"price"`
	Rating      *float64          `json:"rating"`
	BrandID     *uuid.UUID        `json:"brand_id"`
	Description *string           `json:"description"`
	SellerID    *uuid.UUID        `json:"seller_id"`
	Amount      *int32            `json:"amount"`
	Properties  []ProductProperty `json:"properties"`
	Categories  []uuid.UUID       `json:"categories"`
}
