// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: available_passwords.sql

package database

import (
	"context"
)

const insertPassword = `-- name: InsertPassword :one
INSERT INTO available_passwords (password)
VALUES (
    $1
)

RETURNING id, password
`

func (q *Queries) InsertPassword(ctx context.Context, password string) (AvailablePassword, error) {
	row := q.db.QueryRowContext(ctx, insertPassword, password)
	var i AvailablePassword
	err := row.Scan(&i.ID, &i.Password)
	return i, err
}
