-- +goose Up
-- +goose StatementBegin
CREATE TABLE available_passwords (
    id SERIAL PRIMARY KEY,
    password TEXT NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE available_passwords ;
-- +goose StatementEnd
