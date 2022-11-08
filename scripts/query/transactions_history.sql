-- name: CreateTransactionRow :exec
INSERT INTO transactions_history (
    user_id, operation, comments, time, sum
) VALUE (
         ?, ?, ?, ?, ?
    );

-- name: GetTransactionsReport :many
SELECT operation, comments, time, sum
FROM transactions_history
WHERE user_id = ?;