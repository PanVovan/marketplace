// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: product.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (id, name, price, rating, brand_id, seller_id, amount)
VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6) RETURNING id
`

type CreateProductParams struct {
	Name     string        `db:"name"`
	Price    float64       `db:"price"`
	Rating   float64       `db:"rating"`
	BrandID  uuid.NullUUID `db:"brand_id"`
	SellerID uuid.UUID     `db:"seller_id"`
	Amount   int32         `db:"amount"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.Name,
		arg.Price,
		arg.Rating,
		arg.BrandID,
		arg.SellerID,
		arg.Amount,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products WHERE "products"."id" = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

const getProductByID = `-- name: GetProductByID :one
SELECT 
	id,
	name,
	price,
	rating,
	brand_id,
	seller_id,
	amount
FROM "products" AS "product" 
	WHERE "product"."id" = $1
`

func (q *Queries) GetProductByID(ctx context.Context, id uuid.UUID) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProductByID, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.Rating,
		&i.BrandID,
		&i.SellerID,
		&i.Amount,
	)
	return i, err
}
