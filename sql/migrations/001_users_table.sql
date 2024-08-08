-- +goose Up
CREATE TABLE
    users (
        user_id SERIAL PRIMARY KEY,
        email VARCHAR(50) UNIQUE NOT NULL,
        username VARCHAR(50) NOT NULL,
        hashed_password VARCHAR(100) NOT NULL,
        first_name VARCHAR(50) NOT NULL,
        last_name VARCHAR(50) NOT NULL,
        date_of_birth TIMESTAMP NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

-- +goose Down
DROP TABLE users;