-- name: CreateOrder :one
INSERT INTO orders(id, name, email, address, amount, status, phone, comment, user_id) 
VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;

-- name: GetOrderByID :one
SELECT * FROM "orders" AS "order" WHERE id = $1;

-- name: GetOrders :many
SELECT * FROM "orders" AS "order" LIMIT $1 OFFSET $2;

-- name: GetOrdersByUserID :many
SELECT * FROM "orders" AS "order" WHERE user_id = $1 LIMIT $2 OFFSET $3;

-- name: GetOrdersByEmail :many
SELECT * FROM "orders" AS "order" WHERE email = $1 LIMIT $2 OFFSET $3;

-- name: GetOrdersByPhone :many
SELECT * FROM "orders" AS "order" WHERE phone = $1 LIMIT $2 OFFSET $3;

-- name: GetOrdersByStatus :many
SELECT * FROM "orders" AS "order" WHERE status = $1 LIMIT $2 OFFSET $3;

-- name: DeleteOrder :exec
DELETE FROM orders WHERE id = $1;

