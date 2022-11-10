-- name: CreateTransactionRow :exec
INSERT INTO transactions_history (
    user_id, operation, comments, time, sum
) VALUE (
         ?, ?, ?, ?, ?
    );