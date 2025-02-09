-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password TEXT REFERENCES available_passwords(password)  ON DELETE RESTRICT NOT NULL,
    country_code CHAR(2) REFERENCES countries(country_code) ON DELETE RESTRICT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
