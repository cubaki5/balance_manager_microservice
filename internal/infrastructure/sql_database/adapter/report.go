package adapter

import (
	"context"

	"balance_avito/internal/infrastructure/sql_database/generated/sqlc"
	"balance_avito/internal/models"
	"balance_avito/internal/models/types"
)

func (d *DatabaseAdapter) Accounting(ctx context.Context, reportDate models.ReportDate) ([]models.AccountingReport, error) {
	var (
		dbReports []sqlc.GetMonthReportRow
		err       error
	)
	err = d.execTx(ctx, func(q *sqlc.Queries) error {
		dbReports, err = q.GetMonthReport(ctx, sqlc.GetMonthReportParams{
			Year:  reportDate.Year.Int32(),
			Month: reportDate.Month.Int32(),
		})
		return err
	})
	if err != nil {
		return nil, err
	}

	accountringReports := convertDBReportsToModels(dbReports)

	return accountringReports, err
}

func (d *DatabaseAdapter) TransactionsHistory() error {
	//TODO implement me
	panic("implement me")
}

func convertDBReportsToModels(dbReports []sqlc.GetMonthReportRow) []models.AccountingReport {
	var accountingReports = make([]models.AccountingReport, len(dbReports))

	for i := range dbReports {
		accountingReports[i] = models.AccountingReport{
			ServiceID: types.ServiceID(dbReports[i].ServiceID),
			Income:    types.Income(dbReports[i].Income),
		}
	}

	return accountingReports
}
