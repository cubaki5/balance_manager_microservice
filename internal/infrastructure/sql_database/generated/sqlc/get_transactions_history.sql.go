package sqlc

import (
	"context"
	"fmt"
	"time"
)

const getTransactionsReport = `-- name: GetTransactionsReport :many
SELECT operation, comments, time, sum
FROM transactions_history
WHERE user_id = ?
%s
LIMIT ?
OFFSET ?
`

type GetTransactionsReportParams struct {
	SortDate int
	SortSum  int
	UserID   int64
	Limit    int
	Offset   int
}

type GetTransactionsReportRow struct {
	Operation string
	Comments  string
	Time      time.Time
	Sum       int32
}

func makeQueryWithOrder(sortSum int, sortDate int) string {
	var query string

	if sortDate == 1 && sortSum == 1 {
		query = fmt.Sprintf(getTransactionsReport, "ORDER BY time, sum")
	} else if sortDate == 1 {
		query = fmt.Sprintf(getTransactionsReport, "ORDER BY time")
	} else if sortSum == 1 {
		query = fmt.Sprintf(getTransactionsReport, "ORDER BY sum")
	} else {
		query = fmt.Sprintf(getTransactionsReport, "")
	}

	return query
}

func (q *Queries) GetTransactionsReport(ctx context.Context, arg GetTransactionsReportParams) ([]GetTransactionsReportRow, error) {
	query := makeQueryWithOrder(arg.SortSum, arg.SortDate)

	rows, err := q.db.QueryContext(ctx, query, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTransactionsReportRow
	for rows.Next() {
		var i GetTransactionsReportRow
		if err := rows.Scan(
			&i.Operation,
			&i.Comments,
			&i.Time,
			&i.Sum,
		); err != nil {
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
