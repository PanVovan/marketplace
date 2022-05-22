-- name: CreateBrand :one
INSERT INTO brands (id, name) VALUES (uuid_generate_v4(), $1) RETURNING id;
-- name: GetBrands :many
SELECT id, name FROM brands LIMIT $1 OFFSET $2;
-- name: GetBrandById :one
SELECT id, name FROM brands WHERE id = $1;
-- name: DeleteBrand :exec
DELETE FROM brands WHERE brands.id = $1;
