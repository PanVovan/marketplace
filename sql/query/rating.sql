-- name: CreateRating :exec
INSERT INTO rating(id, user_id, product_id, rate) VALUES (uuid_generate_v4(), '63c7948d-24da-4331-9769-ab890505f1bd', '0e1ac0c1-97f9-411e-9657-fabfee9bf65e', 5) RETURNING id;

-- name: GetRatingByProductID :one
SELECT AVG(rate) AS avg FROM rating WHERE product_id = $1;

-- name: DeleteRating :exec
DELETE FROM rating WHERE id = $1;
