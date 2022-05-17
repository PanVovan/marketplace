// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: category.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO categories (id, name) VALUES (uuid_generate_v4(), $1) RETURNING id
`

func (q *Queries) CreateCategory(ctx context.Context, name string) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createCategory, name)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories WHERE "categories"."id" = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const getCategories = `-- name: GetCategories :many
SELECT id, name FROM "categories" as "category"
`

func (q *Queries) GetCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, getCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
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
