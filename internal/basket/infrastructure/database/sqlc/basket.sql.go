// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: basket.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

const createGuestBasket = `-- name: CreateGuestBasket :one
INSERT INTO baskets(id, user_id) VALUES (uuid_generate_v4(), NULL) RETURNING ID
`

func (q *Queries) CreateGuestBasket(ctx context.Context) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createGuestBasket)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const createUserBasket = `-- name: CreateUserBasket :one
INSERT INTO baskets(id, user_id) VALUES (uuid_generate_v4(), $1) RETURNING ID
`

func (q *Queries) CreateUserBasket(ctx context.Context, userID uuid.NullUUID) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createUserBasket, userID)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getBasketByID = `-- name: GetBasketByID :one
SELECT id, user_id FROM baskets WHERE id = $1
`

func (q *Queries) GetBasketByID(ctx context.Context, id uuid.UUID) (Basket, error) {
	row := q.db.QueryRowContext(ctx, getBasketByID, id)
	var i Basket
	err := row.Scan(&i.ID, &i.UserID)
	return i, err
}

const getBasketByUserID = `-- name: GetBasketByUserID :many
SELECT id, user_id FROM baskets WHERE user_id = $1 AND user_id IS NOT NULL
`

func (q *Queries) GetBasketByUserID(ctx context.Context, userID uuid.NullUUID) ([]Basket, error) {
	rows, err := q.db.QueryContext(ctx, getBasketByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Basket
	for rows.Next() {
		var i Basket
		if err := rows.Scan(&i.ID, &i.UserID); err != nil {
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

const getBaskets = `-- name: GetBaskets :many
SELECT id, user_id FROM baskets
`

func (q *Queries) GetBaskets(ctx context.Context) ([]Basket, error) {
	rows, err := q.db.QueryContext(ctx, getBaskets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Basket
	for rows.Next() {
		var i Basket
		if err := rows.Scan(&i.ID, &i.UserID); err != nil {
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
