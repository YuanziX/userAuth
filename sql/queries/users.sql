-- name: CreateUser :one
INSERT INTO
    users (email, username, hashed_password, first_name, last_name, date_of_birth)
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllUsers :many
SELECT *
FROM users;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE email = $1;

-- name: GetHashedPassword :one
SELECT hashed_password
FROM users
WHERE email = $1;

-- name: VerifyUser :exec
UPDATE users
SET verified = TRUE
WHERE email = $1;

-- name: IsUserVerified :one
SELECT verified
FROM users
WHERE email = $1;