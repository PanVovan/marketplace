package aggregate

import (
	brand_dto "makretplace/internal/brand/infrastructure/database/sqlc"
	category_dto "makretplace/internal/category/infrastructure/database/sqlc"
	product_dto "makretplace/internal/product/infrastructure/database/sqlc"
	seller_dto "makretplace/internal/seller/infrastructure/database/sqlc"
)

type Product struct {
	Product    *product_dto.Product
	Images     []*product_dto.ProductImage
	Properties []*product_dto.ProductProperty
	Brand      *brand_dto.Brand
	Categories []*category_dto.Category
	SellerInfo *seller_dto.SellerInfo
}
