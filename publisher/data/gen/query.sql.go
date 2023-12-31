// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package data

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPublisher = `-- name: CreatePublisher :one
INSERT INTO publishers (
    name, bio
) VALUES (
             $1, $2
         )
RETURNING id, name, bio
`

type CreatePublisherParams struct {
	Name string
	Bio  pgtype.Text
}

func (q *Queries) CreatePublisher(ctx context.Context, arg CreatePublisherParams) (Publisher, error) {
	row := q.db.QueryRow(ctx, createPublisher, arg.Name, arg.Bio)
	var i Publisher
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const deletePublisher = `-- name: DeletePublisher :exec
DELETE FROM publishers
WHERE id = $1
`

func (q *Queries) DeletePublisher(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deletePublisher, id)
	return err
}

const getPublisher = `-- name: GetPublisher :one
SELECT id, name, bio FROM publishers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPublisher(ctx context.Context, id int64) (Publisher, error) {
	row := q.db.QueryRow(ctx, getPublisher, id)
	var i Publisher
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const listPublishers = `-- name: ListPublishers :many
SELECT id, name, bio FROM publishers
ORDER BY name
`

func (q *Queries) ListPublishers(ctx context.Context) ([]Publisher, error) {
	rows, err := q.db.Query(ctx, listPublishers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Publisher
	for rows.Next() {
		var i Publisher
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePublisher = `-- name: UpdatePublisher :one
UPDATE publishers
set name = $2,
    bio = $3
WHERE id = $1
    RETURNING id, name, bio
`

type UpdatePublisherParams struct {
	ID   int64
	Name string
	Bio  pgtype.Text
}

func (q *Queries) UpdatePublisher(ctx context.Context, arg UpdatePublisherParams) (Publisher, error) {
	row := q.db.QueryRow(ctx, updatePublisher, arg.ID, arg.Name, arg.Bio)
	var i Publisher
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}
