-- name: GetNames :many
SELECT
    name
FROM
    available_names;

-- name: GetLastNames :many
SELECT
    name
FROM
    available_last;

-- name: GetPasswords :many
SELECT
    password
FROM
    available_passwords;
