// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

// Store the urls created by a user
type Url struct {
	ID        int32
	Url       pgtype.Text
	ShortUrl  pgtype.Text
	Hash      pgtype.Text
	CreatedAt pgtype.Timestamp
	UserID    pgtype.Int4
}

// Users that perform any action in the application
type User struct {
	ID        int32
	Username  pgtype.Text
	Name      pgtype.Text
	Email     pgtype.Text
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}
