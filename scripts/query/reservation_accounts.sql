-- name: DeleteReservedAccount :exec
DELETE FROM reserved_accounts
WHERE order_id = ? LIMIT 1;

-- name: CreateReservedAccount :exec
INSERT INTO reserved_accounts (
    user_id, order_id, service_id, cost
) value (
    ?, ?, ?, ?
    );

-- name: GetReservedAccount :one
SELECT * FROM reserved_accounts
WHERE user_id = ? AND order_id = ? AND service_id = ? AND cost = ? LIMIT 1;