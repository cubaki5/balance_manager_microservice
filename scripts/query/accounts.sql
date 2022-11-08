-- name: ReturnMoney :exec
UPDATE accounts
SET balance = balance + ?
WHERE user_id = ?;

-- name: WriteOffMoney :exec
UPDATE accounts
SET balance = balance - ?
WHERE user_id = ?;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE user_id = ? LIMIT 1;

-- name: CreateAccountOrUpdateBalance :exec
INSERT INTO accounts (
    user_id, balance
) VALUE (
         ?, ?
    )
ON DUPLICATE KEY UPDATE balance = balance + ?;