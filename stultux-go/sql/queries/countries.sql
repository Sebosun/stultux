-- name: InsertCountry :one
INSERT INTO countries (country_code, country_name)
VALUES (
    $1,
    $2
)

RETURNING *;
