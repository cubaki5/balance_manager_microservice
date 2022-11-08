-- name: UpdateAccount :exec
UPDATE accounts
SET balance = ?
WHERE user_id = ?;

-- name: AddAccount :exec
INSERT INTO accounts (
       user_id, balance
) VALUE (
    ?, ?
    );

-- name: GetAccount :one
SELECT * FROM accounts
WHERE user_id = ?