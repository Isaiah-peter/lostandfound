-- name: CreateCategory :one
INSERT INTO categories (
    title,
    discription
) VALUES ($1, $2) RETURNING *;

-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1 LIMIT 1;

-- name: ListCategory :many
SELECT * FROM categories
ORDER BY title
LIMIT $1
OFFSET $2;

-- name: UpdateCategory :exec
UPDATE categories
SET title = $2, discription = $3
WHERE id = $1;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1;