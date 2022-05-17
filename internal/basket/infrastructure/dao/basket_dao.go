package dao

import (
	"context"
	"database/sql"
	"fmt"
	"main/internal/seller/infrastructure/sqlc"
	"strings"

	"github.com/google/uuid"
)

type BasketDao struct {
	db *sql.DB
	*sqlc.Queries
}

func NewBasketDao(db *sql.DB) *BasketDao {
	return &BasketDao{
		db:      db,
		Queries: sqlc.New(db),
	}
}

func (b *BasketDao) UpdateBasket(ctx context.Context, userID uuid.NullUUID, basketID uuid.UUID) error {

	query := `UPDATE baskets SET user_id = $1  WHERE id = $2`

	args := make([]interface{}, 0)
	args = append(args, userID, basketID)

	_, err := b.db.ExecContext(ctx, query, args...)
	return err

}

type UpdateBasketProductParams struct {
	BasketID  *uuid.UUID     `db:"basket_id"`
	ProductID *uuid.UUID     `db:"product_id"`
	Quantity  *sql.NullInt32 `db:"quantity"`
}

func (b *BasketDao) UpdateBasketProduct(ctx context.Context, basketProductID uuid.UUID, arg UpdateBasketProductParams) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if arg.BasketID != nil {
		setValues = append(setValues, fmt.Sprintf("basket_id=$%d", argId))
		args = append(args, *arg.BasketID)
		argId++
	}

	if arg.ProductID != nil {
		setValues = append(setValues, fmt.Sprintf("product_id=$%d", argId))
		args = append(args, *arg.ProductID)
		argId++
	}

	if arg.Quantity != nil {
		setValues = append(setValues, fmt.Sprintf("quantity=$%d", argId))
		args = append(args, *arg.Quantity)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE basket_products SET %s WHERE id = $%d`,
		setQuery, argId)

	args = append(args, basketProductID)

	_, err := b.db.ExecContext(ctx, query, args...)
	return err
}
