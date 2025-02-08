-- +goose Up
-- +goose StatementBegin
CREATE TABLE available_last (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    country_code CHAR(2) REFERENCES countries(country_code) ON DELETE RESTRICT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE available_last;
-- +goose StatementEnd
