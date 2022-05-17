package dao

import (
	"context"
	"database/sql"
	"fmt"
	"main/internal/product/infrastructure/sqlc"
	"strings"

	"github.com/google/uuid"
)

type ProductDao struct {
	db *sql.DB
	*sqlc.Queries
}

func NewProductDao(db *sql.DB) *ProductDao {
	return &ProductDao{
		db:      db,
		Queries: sqlc.New(db),
	}
}

type UpdateProductParams struct {
	Name     *string         `db:"name"`
	Price    *string         `db:"price"`
	Rating   *sql.NullString `db:"rating"`
	BrandID  *uuid.NullUUID  `db:"brand_id"`
	SellerID *uuid.UUID      `db:"seller_id"`
	Amount   *int32          `db:"amount"`
}

func (p *ProductDao) UpdateProduct(ctx context.Context, arg UpdateProductParams, productID uuid.UUID) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

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

	if arg.Rating != nil {
		setValues = append(setValues, fmt.Sprintf("rating=$%d", argId))
		args = append(args, *arg.Rating)
		argId++
	}

	if arg.Amount != nil {
		setValues = append(setValues, fmt.Sprintf("amount=$%d", argId))
		args = append(args, *arg.Amount)
		argId++
	}

	if arg.SellerID != nil {
		setValues = append(setValues, fmt.Sprintf("seller_id=$%d", argId))
		args = append(args, *arg.SellerID)
		argId++
	}

	if arg.BrandID != nil {
		setValues = append(setValues, fmt.Sprintf("brand_id=$%d", argId))
		args = append(args, *arg.BrandID)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE products SET %s WHERE id = $%d`,
		setQuery, argId)

	args = append(args, productID)

	_, err := p.db.ExecContext(ctx, query, args...)
	return err
}

type UpdateProductPropertyParams struct {
	Name  *string `db:"name"`
	Value *string `db:"value"`
}

func (p *ProductDao) UpdateProductProperty(ctx context.Context, arg UpdateProductPropertyParams, productpropertyID uuid.UUID) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if arg.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *arg.Name)
		argId++
	}

	if arg.Value != nil {
		setValues = append(setValues, fmt.Sprintf("value=$%d", argId))
		args = append(args, *arg.Value)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE product_properties SET %s WHERE id = $%d`,
		setQuery, argId)

	args = append(args, productpropertyID)

	_, err := p.db.ExecContext(ctx, query, args...)
	return err
}

func (p *ProductDao) UpdateProductImage(ctx context.Context, file string, imageID uuid.UUID) error {
	query := `UPDATE product_images SET file = $1  WHERE id = $2`

	args := make([]interface{}, 0)
	args = append(args, file, imageID)

	_, err := p.db.ExecContext(ctx, query, args...)
	return err
}
