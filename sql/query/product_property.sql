-- name: CreateProductProperty :one
INSERT INTO product_properties (id, product_id, name, value) VALUES (
	uuid_generate_v4(),
	$1,
	$2,
	$3
) RETURNING id;

-- name: GetProductPropertyByID :one
SELECT * FROM product_properties WHERE id = $1;

-- name: GetProductPropertiesByProductID :many
SELECT * FROM product_properties WHERE product_id = $1;


-- name: DeleteProductProperty :exec
DELETE FROM product_properties WHERE id = $1;

-- name: DeleteProductPropertiesByProductID :exec
DELETE FROM product_properties WHERE product_id = $1;

