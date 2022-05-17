-- name: CreateUser :one
INSERT INTO users(id, email, password, name)
VALUES(uuid_generate_v4(), $1, $2, $3) RETURNING id;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;