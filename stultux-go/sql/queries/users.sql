-- name: CreateUser :one
INSERT INTO users (username, password, country_Code)
VALUES (
    $1,
    $2,
    $3
)

RETURNING *;
