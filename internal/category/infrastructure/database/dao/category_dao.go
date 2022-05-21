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

func (b *CategoryDao) UpdateCategory(ctx context.Context, name string, brandID uuid.UUID) error {
	query := `UPDATE categories SET name = $1  WHERE id = $2`

	args := make([]interface{}, 0)
	args = append(args, name, brandID)

	_, err := b.db.ExecContext(ctx, query, args...)
	return err
}
