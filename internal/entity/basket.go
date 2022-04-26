package entity

type Basket struct {
	ID int `json:"id" db:"id"`
}

type BasketProduct struct {
	BasketID  int `json:"basket_id" db:"basket_id"`
	ProductID int `json:"product_id" db:"product_id"`
	Quantity  int `json:"quantity" db:"quantity"`
}
