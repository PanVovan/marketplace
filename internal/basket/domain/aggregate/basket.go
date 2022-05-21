package aggregate

import (
	basket_dto "makretplace/internal/basket/infrastructure/database/sqlc"
	user_dto "makretplace/internal/user/infrastructure/database/sqlc"
)

type Basket struct {
	Basket   *basket_dto.Basket
	User     *user_dto.UserInfo
	Products []*BasketProduct
}
