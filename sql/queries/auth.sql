-- name: CreateAuth :one
INSERT INTO
    auth (user_email)
VALUES ($1)
    RETURNING *;

-- name: DeleteAuth :exec
DELETE FROM auth
WHERE
    user_email = $1
    AND auth_uuid = $2;

-- name: CheckAuthExists :one
SELECT EXISTS(
    SELECT * FROM auth
    WHERE
        user_email = $1
        AND auth_uuid = $2
);