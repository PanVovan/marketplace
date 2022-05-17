
-- name: CreateCategory :one
INSERT INTO categories (id, name) VALUES (uuid_generate_v4(), $1) RETURNING id;
-- name: GetCategories :many
SELECT id, name FROM "categories" as "category";
-- name: DeleteCategory :exec
DELETE FROM categories WHERE "categories"."id" = $1;