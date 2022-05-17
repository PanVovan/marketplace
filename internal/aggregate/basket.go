package aggregate

import "marketplace/internal/entity"

type Basket struct {
	basket         *entity.Basket
	basketProducts []*entity.BasketProduct
}
