-- name: CreateRating :exec
INSERT INTO rating(id, user_id, product_id, rate) VALUES (uuid_generate_v4(), $1, $2, $3) RETURNING id;

-- name: GetRatingByProductID :one
SELECT AVG(rate) AS avg FROM rating WHERE product_id = $1;

-- name: DeleteRating :exec
DELETE FROM rating WHERE id = $1;
