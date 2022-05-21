package dao

import (
	"context"
	"database/sql"
	"makretplace/internal/brand/infrastructure/database/sqlc"

	"github.com/google/uuid"
)

type BrandDao struct {
	db *sql.DB
	*sqlc.Queries
}

func NewBrandDao(db *sql.DB) *BrandDao {
	return &BrandDao{
		db:      db,
		Queries: sqlc.New(db),
	}
}

func (b *BrandDao) UpdateBrand(ctx context.Context, name string, brandID uuid.UUID) error {
	query := `UPDATE brands SET name = $1  WHERE id = $2`

	args := make([]interface{}, 0)
	args = append(args, name, brandID)

	_, err := b.db.ExecContext(ctx, query, args...)
	return err
}
