package dao

import (
	"context"
	"database/sql"
	"fmt"
	"makretplace/internal/seller/infrastructure/database/sqlc"
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

const getSellerInfoByID = `
SELECT id, name FROM "sellers" as "seller" WHERE "seller"."id" = $1
`

func (q *SellerDao) GetSellerInfoByID(ctx context.Context, id uuid.UUID) (sqlc.SellerInfo, error) {
	row := q.db.QueryRowContext(ctx, getSellerInfoByID, id)
	var i sqlc.SellerInfo
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getSellersInfo = `
SELECT id, name FROM "sellers" as "seller" LIMIT $1 OFFSET $2
`

type GetSellersInfoParams struct {
	Limit  int32 `db:"limit"`
	Offset int32 `db:"offset"`
}

func (q *SellerDao) GetSellersInfo(ctx context.Context, arg GetSellersInfoParams) ([]sqlc.SellerInfo, error) {
	rows, err := q.db.QueryContext(ctx, getSellersInfo, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sqlc.SellerInfo
	for rows.Next() {
		var i sqlc.SellerInfo
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
