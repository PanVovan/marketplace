-- name: CreateImage :one
INSERT INTO "product_images" (id, product_id, file) VALUES (
	uuid_generate_v4(),
	$1,
	$2
) RETURNING id;

-- name: GetImages :many
SELECT id, product_id, file FROM product_images;

-- name: GetImageByID :one
SELECT id, product_id, file FROM product_images WHERE id = $1;
-- name: GetImagesByProductID :many
SELECT id, product_id, file FROM product_images WHERE product_id = $1;


-- name: DeleteImage :exec
DELETE FROM product_images WHERE id = $1;

-- name: DeleteImagesByProductID :exec
DELETE FROM product_images WHERE product_id = $1;

