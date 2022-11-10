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

	accountringReports := convertDBReportToModels(dbReports)

	return accountringReports, err
}

func (d *DatabaseAdapter) TransactionsHistory(ctx context.Context, transHistoryParams models.TransactionsHistoryParams) ([]models.TransactionsHistory, error) {
	var (
		dbReports []sqlc.GetTransactionsReportRow
		err       error
	)
	err = d.execTx(ctx, func(q *sqlc.Queries) error {
		dbReports, err = q.GetTransactionsReport(ctx, sqlc.GetTransactionsReportParams{
			SortDate: transHistoryParams.SortDate,
			SortSum:  transHistoryParams.SortSum,
			UserID:   transHistoryParams.UserID.Int64(),
			Limit:    transHistoryParams.Limit,
			Offset:   (transHistoryParams.Page - 1) * transHistoryParams.Limit,
		})
		return err
	})
	if err != nil {
		return nil, err
	}

	accountringReports := convertDBHistoryToModels(dbReports)

	return accountringReports, err
}

func convertDBReportToModels(dbReports []sqlc.GetMonthReportRow) []models.AccountingReport {
	var accountingReports = make([]models.AccountingReport, len(dbReports))

	for i := range dbReports {
		accountingReports[i] = models.AccountingReport{
			ServiceID: types.ServiceID(dbReports[i].ServiceID),
			Income:    types.Income(dbReports[i].Income),
		}
	}

	return accountingReports
}

func convertDBHistoryToModels(dbHistory []sqlc.GetTransactionsReportRow) []models.TransactionsHistory {
	var transactionHistory = make([]models.TransactionsHistory, len(dbHistory))

	for i := range dbHistory {
		transactionHistory[i].Operation = types.Operation(dbHistory[i].Operation)
		transactionHistory[i].Sum = types.Income(dbHistory[i].Sum)
		transactionHistory[i].Time = dbHistory[i].Time
		transactionHistory[i].Comments = types.Comment(dbHistory[i].Comments)
	}

	return transactionHistory
}
