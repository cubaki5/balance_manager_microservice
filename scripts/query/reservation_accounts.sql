-- name: DeleteReservedAccount :exec
DELETE FROM reserved_accounts
WHERE order_id = ?;

-- name: CreateReservedAccount :exec
INSERT INTO reserved_accounts (
    user_id, order_id, service_id, cost
) value (
    ?, ?, ?, ?
    );