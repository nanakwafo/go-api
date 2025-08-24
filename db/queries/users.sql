-- name: CreateUser :one
INSERT INTO users (name, email)
VALUES ($1, $2)
RETURNING id, name, email;

-- name: GetUser :one
SELECT id, name, email FROM users WHERE id = $1;

-- name: GetUsers :many
SELECT id, name, email FROM users;

-- name: UpdateUser :exec
UPDATE users SET name = $1, email = $2 WHERE id = $3;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
