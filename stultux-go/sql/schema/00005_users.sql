-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT REFERENCES available_names(name) ON DELETE RESTRICT,
    last_name TEXT REFERENCES available_last(name) ON DELETE RESTRICT,
    password TEXT REFERENCES available_passwords(password)  ON DELETE RESTRICT,
    country_code CHAR(2) REFERENCES countries(country_code) ON DELETE RESTRICT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
