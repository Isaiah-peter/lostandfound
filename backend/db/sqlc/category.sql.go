// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: category.sql

package db

import (
	"context"
	"database/sql"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO categories (
    title,
    discription
) VALUES ($1, $2) RETURNING id, title, discription, created_at
`

type CreateCategoryParams struct {
	Title       sql.NullString `json:"title"`
	Discription sql.NullString `json:"discription"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, createCategory, arg.Title, arg.Discription)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Discription,
		&i.CreatedAt,
	)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const getCategory = `-- name: GetCategory :one
SELECT id, title, discription, created_at FROM categories
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCategory(ctx context.Context, id int32) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategory, id)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Discription,
		&i.CreatedAt,
	)
	return i, err
}

const listCategory = `-- name: ListCategory :many
SELECT id, title, discription, created_at FROM categories
ORDER BY title
LIMIT $1
OFFSET $2
`

type ListCategoryParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCategory(ctx context.Context, arg ListCategoryParams) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, listCategory, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Discription,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCategory = `-- name: UpdateCategory :exec
UPDATE categories
SET title = $2, discription = $3
WHERE id = $1
`

type UpdateCategoryParams struct {
	ID          int32          `json:"id"`
	Title       sql.NullString `json:"title"`
	Discription sql.NullString `json:"discription"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, updateCategory, arg.ID, arg.Title, arg.Discription)
	return err
}
