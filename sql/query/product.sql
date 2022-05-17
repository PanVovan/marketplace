
-- name: CreateProduct :one
INSERT INTO products (id, name, price, rating, brand_id, seller_id, amount)
VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6) RETURNING id;

-- name: GetProductByID :one
SELECT 
	id,
	name,
	price,
	rating,
	brand_id,
	seller_id,
	amount
FROM "products" AS "product" 
	WHERE "product"."id" = $1;

-- name: GetProducts :many
SELECT 
	id,
	name,
	price,
	rating,
	brand_id, 
	seller_id,
	amount
FROM "products" AS "product" LIMIT $1 OFFSET $2;

-- name: GetProductsBySellerID :many
SELECT 
	id,
	name,
	price,
	rating,
	brand_id, 
	seller_id,
	amount
FROM "products" AS "product" WHERE product.seller_id = $1 LIMIT $2 OFFSET $3; 

-- name: DeleteProduct :exec
DELETE FROM products WHERE "products"."id" = $1;
