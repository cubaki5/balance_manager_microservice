-- name: CreateOrUpdateReportRow :exec
INSERT INTO report_accounting (
    service_id, year, month, income
) VALUE (
    ?, ?, ?, ?
    )
ON DUPLICATE KEY UPDATE income = income + ?;

-- name: GetMonthReport :many
SELECT service_id, income
FROM report_accounting
WHERE year = ? AND month = ?;