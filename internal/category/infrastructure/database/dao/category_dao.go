package dao

import (
	"context"
	"database/sql"
	"makretplace/internal/category/infrastructure/database/sqlc"

	"github.com/google/uuid"
)

type CategoryDao struct {
	db *sql.DB
	*sqlc.Queries
}

func NewCategoryDao(db *sql.DB) *CategoryDao {
	return &CategoryDao{
		db:      db,
		Queries: sqlc.New(db),
	}

}

func (b *CategoryDao) UpdateCategory(ctx context.Context, categoryId uuid.UUID, name string) error {
	query := `UPDATE categories SET name = $1  WHERE id = $2`

	args := make([]interface{}, 0)
	args = append(args, name, categoryId)

	_, err := b.db.ExecContext(ctx, query, args...)
	return err
}

const deleteProductCategoryByCategoryID = `-- name: DeleteProductCategoryByCategoryID :exec
DELETE FROM products_categories WHERE categories_id = $1
`

func (q *CategoryDao) DeleteProductCategoryByCategoryID(ctx context.Context, categoriesID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteProductCategoryByCategoryID, categoriesID)
	return err
}
