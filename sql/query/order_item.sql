-- name: CreateOrderItem :one
INSERT INTO order_items(id, order_id, name, price, quantity)
VALUES (uuid_generate_v4(), $1, $2, $3, $4) RETURNING id;

-- name: GetOrderItems :many
SELECT * FROM order_items LIMIT $1 OFFSET $2;

-- name: GetOrderItemsByOrderID :many
SELECT * FROM order_items WHERE order_id = $1 LIMIT $2 OFFSET $3;

-- name: DeleteOrderItem :exec
DELETE FROM order_items WHERE id = $1;