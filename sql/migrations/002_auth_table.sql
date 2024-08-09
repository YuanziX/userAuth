-- +goose Up
CREATE TABLE
    auth (
        auth_id SERIAL PRIMARY KEY,
        user_email VARCHAR(50) NOT NULL,
        auth_uuid UUID NOT NULL DEFAULT gen_random_uuid ()
    );

-- +goose Down
DROP TABLE auth;