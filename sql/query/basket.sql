-- name: CreateUserBasket :one
INSERT INTO baskets(id, user_id) VALUES (uuid_generate_v4(), $1) RETURNING ID;

-- name: CreateGuestBasket :one
INSERT INTO baskets(id, user_id) VALUES (uuid_generate_v4(), NULL) RETURNING ID;

-- name: GetBaskets :many
SELECT * FROM baskets ;

-- name: GetBasketByID :one
SELECT * FROM baskets WHERE id = $1;

-- name: GetBasketByUserID :many
SELECT * FROM baskets WHERE user_id = $1 AND user_id IS NOT NULL;