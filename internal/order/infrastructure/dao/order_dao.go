package dao

import (
	"context"
	"database/sql"
	"fmt"
	"main/internal/order/infrastructure/sqlc"
	"strings"

	"github.com/google/uuid"
)

type OrderDao struct {
	*sqlc.Queries
	db *sql.DB
}

type UpdateOrderParams struct {
	Name    *string         `db:"name"`
	Email   *string         `db:"email"`
	Address *string         `db:"address"`
	Amount  *int32          `db:"amount"`
	Status  *int32          `db:"status"`
	Phone   *string         `db:"phone"`
	Comment *sql.NullString `db:"comment"`
	UserID  *uuid.NullUUID  `db:"user_id"`
}

func NewOrderDao(db *sql.DB) *OrderDao {
	return &OrderDao{
		db:      db,
		Queries: sqlc.New(db),
	}
}

func (o *OrderDao) UpdateOrder(ctx context.Context, orderId uuid.UUID, arg UpdateOrderParams) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if arg.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *arg.Email)
		argId++
	}

	if arg.Address != nil {
		setValues = append(setValues, fmt.Sprintf("address=$%d", argId))
		args = append(args, *arg.Address)
		argId++
	}

	if arg.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *arg.Name)
		argId++
	}

	if arg.Amount != nil {
		setValues = append(setValues, fmt.Sprintf("amount=$%d", argId))
		args = append(args, *arg.Amount)
		argId++
	}

	if arg.Status != nil {
		setValues = append(setValues, fmt.Sprintf("status=$%d", argId))
		args = append(args, *arg.Status)
		argId++
	}

	if arg.UserID != nil {
		setValues = append(setValues, fmt.Sprintf("user_id=$%d", argId))
		args = append(args, *arg.Phone)
		argId++
	}

	if arg.Comment != nil {
		setValues = append(setValues, fmt.Sprintf("comment=$%d", argId))
		args = append(args, *arg.Status)
		argId++
	}

	if arg.Phone != nil {
		setValues = append(setValues, fmt.Sprintf("phone=$%d", argId))
		args = append(args, *arg.Phone)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE orders SET %s WHERE id = $%d`,
		setQuery, argId)

	args = append(args, orderId)

	_, err := o.db.ExecContext(ctx, query, args...)
	return err
}

type UpdateOrderItemParams struct {
	OrderID  *uuid.UUID `db:"order_id"`
	Name     *string    `db:"name"`
	Price    *float64   `db:"price"`
	Quantity *int32     `db:"quantity"`
}

func (o *OrderDao) UpdateOrderItem(ctx context.Context, orderId uuid.UUID, arg UpdateOrderItemParams) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if arg.OrderID != nil {
		setValues = append(setValues, fmt.Sprintf("order_id=$%d", argId))
		args = append(args, *arg.OrderID)
		argId++
	}

	if arg.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *arg.Name)
		argId++
	}

	if arg.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *arg.Price)
		argId++
	}

	if arg.Quantity != nil {
		setValues = append(setValues, fmt.Sprintf("quantity=$%d", argId))
		args = append(args, *arg.Quantity)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE order_items SET %s WHERE id = $%d`,
		setQuery, argId)

	args = append(args, orderId)

	_, err := o.db.ExecContext(ctx, query, args...)
	return err
}
