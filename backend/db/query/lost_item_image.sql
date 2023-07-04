-- name: CreateImage :one
INSERT INTO lost_items_images (
    lost_item_id,
    lost_item_image
) VALUES ($1, $2) RETURNING *;

-- name: GetImage :one
SELECT * FROM lost_items_images
WHERE id = $1 LIMIT 1;

-- name: GetImageByLID :one
SELECT * FROM lost_items_images
WHERE lost_item_id = $1 LIMIT 5;

-- name: DeleteImage :exec
DELETE FROM lost_items_images
WHERE id = $1;

-- name: DeleteImageByLID :exec
DELETE FROM lost_items_images
WHERE lost_item_id = $1;

