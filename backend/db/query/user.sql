-- name: CreateUser :one
INSERT INTO users (
    full_name,
    address,
    contact,
    username,
    user_image,
    password
) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByUserName :one
SELECT * FROM users
WHERE full_name = $1 LIMIT 1;
