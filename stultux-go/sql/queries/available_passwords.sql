-- name: InsertPassword :one
INSERT INTO available_passwords (password)
VALUES (
    $1
)

RETURNING *;
