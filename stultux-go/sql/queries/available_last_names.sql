-- name: InsertLastName :one
INSERT INTO available_last (name, country_code)
VALUES (
    $1,
    $2
)

RETURNING *;
