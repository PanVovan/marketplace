package entity

import "github.com/google/uuid"

// Order model, DB table: "orders"
// One user can have many orders
type Order struct {
	ID      uuid.UUID `json:"-" db:"id"`
	UserId  uuid.UUID `json:"user_id" db:"user_id"`
	Email   string    `json:"email" db:"email"`
	Name    string    `json:"name" db:"name"`
	Phone   string    `json:"phone" db:"phone"`
	Address string    `json:"address" db:"address"`
	Amount  string    `json:"amount" db:"amount"`
	Status  string    `json:"status" db:"status"`
	Comment string    `json:"comment" db:"comment"`
}

// Order item model, DB table: "items"
// One order can have many items

type OrderItem struct {
	ID       uuid.UUID `json:"-" db:"id"`
	OrderID  uuid.UUID `json:"-" db:"order_id"`
	Name     string    `json:"name" db:"name"`
	Price    int       `json:"price" db:"price"`
	Quantity int       `json:"quantity" db:"quantity"`
}
