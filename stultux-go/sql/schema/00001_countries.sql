-- apparentlt there need to be two tabs not one in here or it breaks --.-.-.-.-
-- +goose Up
-- +goose StatementBegin
CREATE TABLE countries (
    country_code CHAR(2) PRIMARY KEY, -- ISO country code (e.g., 'US', 'IN', 'CA')
    country_name TEXT NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE countries;
-- +goose StatementEnd
