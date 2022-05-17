package dao

import (
	"context"
	"database/sql"
	"fmt"
	"main/internal/seller/infrastructure/sqlc"
	"strings"

	"github.com/google/uuid"
)

type SellerDao struct {
	*sqlc.Queries
	db *sql.DB
}

type UpdateSellerParams struct {
	Email    *string `db:"email"`
	Password *string `db:"password"`
	Name     *string `db:"name"`
}

func NewSellerDao(db *sql.DB) *SellerDao {
	return &SellerDao{
		db:      db,
		Queries: sqlc.New(db),
	}
}

func (s *SellerDao) UpdateSeller(ctx context.Context, sellerId uuid.UUID, arg UpdateSellerParams) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if arg.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *arg.Email)
		argId++
	}

	if arg.Password != nil {
		setValues = append(setValues, fmt.Sprintf("password=$%d", argId))
		args = append(args, *arg.Password)
		argId++
	}

	if arg.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *arg.Name)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE sellers SET %s WHERE id = $%d`,
		setQuery, argId)

	args = append(args, sellerId)
	_, err := s.db.ExecContext(ctx, query, args...)
	return err
}
