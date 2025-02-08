-- name: InsertName :one
INSERT INTO available_names (name, country_code)
VALUES (
    $1,
    $2
)

RETURNING *;
