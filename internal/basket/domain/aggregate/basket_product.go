package aggregate

import (
	basket_product_dto "makretplace/internal/basket/infrastructure/database/sqlc"
	"makretplace/internal/product/domain/aggregate"
	image_dto "makretplace/internal/product/infrastructure/database/sqlc"
)

type BasketProduct struct {
	Product       *aggregate.ProductInfo
	BasketProduct *basket_product_dto.BasketProduct
	Image         *image_dto.ProductImage
}
