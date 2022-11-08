-- name: CreateReportRow :exec
INSERT INTO report_accounting (
    service_id, date, income
) VALUE (
    ?, ?, ?
    );

-- name: GetMonthReport :many
SELECT service_id, income
FROM report_accounting
WHERE YEAR(date) = ? AND MONTH(date) = ?;