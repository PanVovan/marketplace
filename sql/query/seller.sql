CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- name: CreateSeller :one
INSERT INTO sellers (id, name, email, password)
VALUES (uuid_generate_v4(), $1, $2, $3) RETURNING id;
-- name: GetSellerByEmail :one
SELECT id, name, email, password FROM sellers as seller WHERE seller.email = $1;
-- name: GetSellerByID :one
SELECT id, name, email, password FROM sellers as seller WHERE seller.id = $1;
-- name: GetSellers :many
SELECT id, name, email, password FROM sellers as seller LIMIT $1 OFFSET $2;
-- name: DeleteSeller :exec
DELETE FROM sellers WHERE sellers.id = $1;
