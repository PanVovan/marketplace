package entity

import "github.com/google/uuid"

type Basket struct {
	ID uuid.UUID `json:"id" db:"id"`
}

type BasketProduct struct {
	BasketID  uuid.UUID `json:"basket_id" db:"basket_id"`
	ProductID uuid.UUID `json:"product_id" db:"product_id"`
	Quantity  int       `json:"quantity" db:"quantity"`
}
