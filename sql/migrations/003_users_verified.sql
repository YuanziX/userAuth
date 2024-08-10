-- +goose Up
ALTER TABLE users ADD verified BOOLEAN NOT NULL DEFAULT FALSE;

-- +goose Down
ALTER TABLE users
DROP COLUMN verified;