-- name: CreateBasketProduct :one
INSERT INTO "basket_products"(id, basket_id, product_id, quantity) VALUES(uuid_generate_v4(), $1, $2, $3) RETURNING id;

-- name: GetBasketProducts :many
SELECT * FROM "basket_products" LIMIT $1 OFFSET $2;

-- name: GetBasketProductByID :one
SELECT id, basket_id, product_id, quantity FROM "basket_products" WHERE id = $1;

-- name: GetBasketProductByProductID :one
SELECT id, basket_id, product_id, quantity FROM "basket_products" WHERE product_id = $1;

-- name: GetBasketProductsByBasketID :many
 SELECT id, basket_id, product_id, quantity FROM "basket_products" WHERE basket_id = $1 LIMIT $2 OFFSET $3;

-- name: DeleteBasketProduct :exec
DELETE FROM basket_products WHERE id = $1;

-- name: DeleteBasketProductByBasketID :exec
DELETE FROM basket_products WHERE basket_id = $1;

-- name: DeleteBasketProductByProductID :exec
DELETE FROM basket_products WHERE product_id = $1;