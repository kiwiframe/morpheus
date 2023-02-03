-- name: GetAccount :one
SELECT *
FROM users.accounts
WHERE id = $1;
