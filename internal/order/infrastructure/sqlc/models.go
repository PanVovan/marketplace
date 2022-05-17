// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package sqlc

import (
	"database/sql"

	"github.com/google/uuid"
)

type Order struct {
	ID      uuid.UUID      `db:"id"`
	Name    string         `db:"name"`
	Email   string         `db:"email"`
	Address string         `db:"address"`
	Amount  int32          `db:"amount"`
	Status  int32          `db:"status"`
	Phone   string         `db:"phone"`
	Comment sql.NullString `db:"comment"`
	UserID  uuid.NullUUID  `db:"user_id"`
}

type OrderItem struct {
	ID       uuid.UUID `db:"id"`
	OrderID  uuid.UUID `db:"order_id"`
	Name     string    `db:"name"`
	Price    float64   `db:"price"`
	Quantity int32     `db:"quantity"`
}