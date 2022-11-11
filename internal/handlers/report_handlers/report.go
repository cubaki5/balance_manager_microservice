package report_handlers

import (
	"context"

	"balance_avito/internal/models"
)

//go:generate mockgen --source=report.go --destination=mocks/mock_report.go --package=mocks

type (
	GoCSVAdapter interface {
		MarshalStructHistoryToCSV(transactionsHistory []models.TransactionsHistory) (string, error)
		MarshalStructReportToCSV(report []models.AccountingReport) (string, error)
	}

	Database interface {
		Accounting(ctx context.Context, reportDate models.ReportDate) ([]models.AccountingReport, error)
		TransactionsHistory(ctx context.Context, transactionsHistoryParams models.TransactionsHistoryParams) ([]models.TransactionsHistory, error)
	}
)

type Report struct {
	db    Database
	gocsv GoCSVAdapter
}

func NewReport(db Database, gocsv GoCSVAdapter) *Report {
	return &Report{db, gocsv}
}

func (r *Report) Accounting(ctx context.Context, reportDate models.ReportDate) (string, error) {
	err := reportDate.Validate()
	if err != nil {
		return "", err
	}

	accountReports, err := r.db.Accounting(ctx, reportDate)
	if err != nil {
		return "", err
	}

	report, err := r.gocsv.MarshalStructReportToCSV(accountReports)
	if err != nil {
		return "", err
	}

	return report, nil
}

func (r *Report) TransactionsHistory(ctx context.Context, transactionsHistoryParams models.TransactionsHistoryParams) (string, error) {
	err := transactionsHistoryParams.Validate()
	if err != nil {
		return "", err
	}

	transactionsHistory, err := r.db.TransactionsHistory(ctx, transactionsHistoryParams)
	if err != nil {
		return "", err
	}

	transactionsHistoryCSV, err := r.gocsv.MarshalStructHistoryToCSV(transactionsHistory)
	if err != nil {
		return "", err
	}

	return transactionsHistoryCSV, nil
}
