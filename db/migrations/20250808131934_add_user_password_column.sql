-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN password VARCHAR NOT NULL DEFAULT 'password';
-- +goose StatementEnd

-- +goose Down
ALTER TABLE users
DROP COLUMN IF EXISTS password;
-- +goose StatementBegin
-- +goose StatementEnd
