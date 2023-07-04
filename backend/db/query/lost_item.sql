-- name: CreateLostItem :one
INSERT INTO lost_items (
    category_id,
    founder_id,
    title,
    discription,
    date,
    time,
    location,
    post_type,
    status,
    remark
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING *;

-- name: GetLostItem :one
SELECT * FROM lost_items
WHERE id = $1 LIMIT 1;

-- name: ListLostItem :many
SELECT * FROM lost_items
LIMIT $1
OFFSET $2;


-- name: UpdateLostItemStatus :exec
UPDATE lost_items
SET status = $2
WHERE id = $1;

-- name: DeleteLostItem :exec
DELETE FROM lost_items
WHERE id = $1;
