-- name: CreateUser :one
INSERT INTO users (name, last_name, password, country_Code)
VALUES (
    $1,
    $2,
    $3,
    $4
)

RETURNING *;
