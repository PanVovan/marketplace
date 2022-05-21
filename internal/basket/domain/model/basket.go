package model

import (
	user_model "makretplace/internal/user/domain/model"

	"github.com/google/uuid"
)

type Basket struct {
	ID             uuid.UUID
	User           user_model.UserInfo
	BasketProducts []BasketProduct
}
