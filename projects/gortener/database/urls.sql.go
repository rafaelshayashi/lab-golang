// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: urls.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createURL = `-- name: CreateURL :one
INSERT INTO urls (
  url, short_url, hash, user_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING id
`

type CreateURLParams struct {
	Url      pgtype.Text
	ShortUrl pgtype.Text
	Hash     pgtype.Text
	UserID   pgtype.Int4
}

func (q *Queries) CreateURL(ctx context.Context, arg CreateURLParams) (int32, error) {
	row := q.db.QueryRow(ctx, createURL,
		arg.Url,
		arg.ShortUrl,
		arg.Hash,
		arg.UserID,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getURL = `-- name: GetURL :one
SELECT id, url, short_url, hash, created_at, user_id FROM urls
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetURL(ctx context.Context, id int32) (Url, error) {
	row := q.db.QueryRow(ctx, getURL, id)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.ShortUrl,
		&i.Hash,
		&i.CreatedAt,
		&i.UserID,
	)
	return i, err
}