// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: report_accounting.sql

package sqlc

import (
	"context"
)

const createOrUpdateReportRow = `-- name: CreateOrUpdateReportRow :exec
INSERT INTO report_accounting (
    service_id, year, month, income
) VALUE (
    ?, ?, ?, ?
    )
ON DUPLICATE KEY UPDATE income = income + ?
`

type CreateOrUpdateReportRowParams struct {
	ServiceID int64
	Year      int32
	Month     int32
	Income    int32
	Income_2  int32
}

func (q *Queries) CreateOrUpdateReportRow(ctx context.Context, arg CreateOrUpdateReportRowParams) error {
	_, err := q.db.ExecContext(ctx, createOrUpdateReportRow,
		arg.ServiceID,
		arg.Year,
		arg.Month,
		arg.Income,
		arg.Income_2,
	)
	return err
}

const getMonthReport = `-- name: GetMonthReport :many
SELECT service_id, income
FROM report_accounting
WHERE year = ? AND month = ?
`

type GetMonthReportParams struct {
	Year  int32
	Month int32
}

type GetMonthReportRow struct {
	ServiceID int64
	Income    int32
}

func (q *Queries) GetMonthReport(ctx context.Context, arg GetMonthReportParams) ([]GetMonthReportRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonthReport, arg.Year, arg.Month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMonthReportRow
	for rows.Next() {
		var i GetMonthReportRow
		if err := rows.Scan(&i.ServiceID, &i.Income); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
