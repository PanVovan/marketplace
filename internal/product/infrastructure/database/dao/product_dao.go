package dao

import (
	"context"
	"database/sql"
	"fmt"
	category_dto "makretplace/internal/category/infrastructure/database/sqlc"
	"makretplace/internal/product/infrastructure/database/sqlc"
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
	Name        *string    `db:"name"`
	Price       *float64   `db:"price"`
	Rating      *float64   `db:"rating"`
	Description *string    `db:"description"`
	BrandID     *uuid.UUID `db:"brand_id"`
	SellerID    *uuid.UUID `db:"seller_id"`
	Amount      *int32     `db:"amount"`
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

	if arg.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *arg.Description)
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

func (p *ProductDao) UpdateProductProperty(ctx context.Context, productpropertyID uuid.UUID, arg UpdateProductPropertyParams) error {
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

type GetProductsParams struct {
	Limit  int32 `db:"limit"`
	Offset int32 `db:"offset"`
}

type GetProductsQuerySpecs struct {
	SellerID     *uuid.UUID
	CategoriesID []*uuid.UUID
	BrandID      *uuid.UUID
}

func (p *ProductDao) GetProducts(ctx context.Context, arg GetProductsParams, specs GetProductsQuerySpecs) ([]sqlc.Product, error) {
	var values []string = nil
	args := make([]interface{}, 0)
	argId := 1

	join := ""
	where := ""

	if specs.BrandID != nil {

		values = append(values, fmt.Sprintf("brand_id = $%d", argId))
		args = append(args, *specs.BrandID)
		argId++

	}

	if specs.SellerID != nil {

		values = append(values, fmt.Sprintf("seller_id = $%d", argId))
		args = append(args, *specs.SellerID)
		argId++

	}

	if specs.CategoriesID != nil {
		join = `
		INNER JOIN products_categories ON products.id = products_categories.products_id
		INNER JOIN categories ON categories.id = products_categories.categories_id`

		var category []string = nil

		for _, categoryID := range specs.CategoriesID {
			category = append(category, fmt.Sprintf("categories.id = $%d", argId))
			args = append(args, *categoryID)
			argId++
		}

		if category != nil {
			values = append(values, strings.Join(category, " OR "))
		}

	}

	if len(values) != 0 {
		where = fmt.Sprintf("WHERE %s", strings.Join(values, " AND "))
	}

	query := fmt.Sprintf(`
	SELECT 
		products.id,
		products.name,
		products.description
		products.price,
		products.rating,
		products.brand_id,
		products.seller_id,
		products.amount
	FROM products %s 
	%s LIMIT $%d OFFSET $%d
	`,
		join, where, argId, argId+1,
	)

	args = append(args, arg.Limit, arg.Offset)

	rows, err := p.db.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sqlc.Product
	for rows.Next() {
		var i sqlc.Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.Rating,
			&i.BrandID,
			&i.SellerID,
			&i.Amount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil

}

const getCategoriesByProductID = `-- name: GetCategoriesByProductID :many
SELECT id, name FROM categories INNER JOIN products_categories ON categories.id = products_categories.categories_id WHERE products_categories.products_id = $1
`

func (q *ProductDao) GetCategoriesByProductID(ctx context.Context, productsID uuid.UUID) ([]category_dto.Category, error) {
	rows, err := q.db.QueryContext(ctx, getCategoriesByProductID, productsID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []category_dto.Category
	for rows.Next() {
		var i category_dto.Category
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
