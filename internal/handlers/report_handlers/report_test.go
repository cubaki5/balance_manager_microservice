package report_handlers

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"balance_avito/internal/handlers/report_handlers/mocks"
	"balance_avito/internal/models"
)

func initDatabaseMock(t *testing.T) *mocks.MockDatabase {
	ctrl := gomock.NewController(t)
	return mocks.NewMockDatabase(ctrl)
}

func initMarshalCSVMock(t *testing.T) *mocks.MockGoCSVAdapter {
	ctrl := gomock.NewController(t)
	return mocks.NewMockGoCSVAdapter(ctrl)
}

func TestReport_Accounting(t *testing.T) {
	tests := []struct {
		name        string
		reportDate  models.ReportDate
		expErr      error
		expDBReport []models.AccountingReport
		expReport   string
		assertErr   func(t *testing.T, expErr error, actErr error)
	}{
		{
			name: "Happy Path",
			reportDate: models.ReportDate{
				Year:  2022,
				Month: 11,
			},
			expErr: nil,
			expDBReport: []models.AccountingReport{
				{ServiceID: 1,
					Income: 1,
				},
			},
			expReport: "testReport",
			assertErr: func(t *testing.T, _ error, actErr error) {
				assert.NoError(t, actErr)
			},
		},
		{
			name: "Wrong Year",
			reportDate: models.ReportDate{
				Year:  1810,
				Month: 11,
			},
			expErr:    errors.New("неверное значение year"),
			expReport: "",
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
		{
			name: "Negative Month",
			reportDate: models.ReportDate{
				Year:  1810,
				Month: -11,
			},
			expErr:    errors.New("неверное значение month"),
			expReport: "",
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
		{
			name: "Wrong Month",
			reportDate: models.ReportDate{
				Year:  1810,
				Month: 13,
			},
			expErr:    errors.New("неверное значение month"),
			expReport: "",
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
		{
			name: "Accounting returns err",
			reportDate: models.ReportDate{
				Year:  2022,
				Month: 11,
			},
			expErr: errors.New("testErr"),
			expDBReport: []models.AccountingReport{
				{},
			},
			expReport: "",
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
		{
			name: "Marshaling returns err",
			reportDate: models.ReportDate{
				Year:  2022,
				Month: 11,
			},
			expErr: errors.New("testErr"),
			expDBReport: []models.AccountingReport{
				{ServiceID: 1,
					Income: 1,
				},
			},
			expReport: "",
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockDB := initDatabaseMock(t)
			mockMarshaling := initMarshalCSVMock(t)
			r := NewReport(mockDB, mockMarshaling)

			ctx := context.Background()
			if test.name == "Happy Path" || test.name == "Marshaling returns err" {
				mockDB.EXPECT().Accounting(ctx, test.reportDate).Return(test.expDBReport, nil)
				mockMarshaling.EXPECT().MarshalStructReportToCSV(test.expDBReport).Return(test.expReport, test.expErr)
			}
			if test.name == "Accounting returns err" {
				mockDB.EXPECT().Accounting(ctx, test.reportDate).Return(test.expDBReport, test.expErr)
			}

			actReport, actErr := r.Accounting(ctx, test.reportDate)

			test.assertErr(t, test.expErr, actErr)
			assert.Equal(t, test.expReport, actReport)
		})
	}
}

func TestReport_TransactionsHistory(t *testing.T) {
	tests := []struct {
		name         string
		params       models.TransactionsHistoryParams
		expErr       error
		expDBHistory []models.TransactionsHistory
		expReport    string
		assertErr    func(t *testing.T, expErr error, actErr error)
	}{
		{
			name: "Happy Path",
			params: models.TransactionsHistoryParams{
				UserID:   1,
				SortDate: 1,
				SortSum:  1,
				Page:     1,
				Limit:    1,
			},
			expErr: nil,
			expDBHistory: []models.TransactionsHistory{
				{Operation: "testOperation",
					Comments: "testComment",
					Time:     time.Unix(0, 0),
					Sum:      1,
				},
			},
			expReport: "testReport",
			assertErr: func(t *testing.T, _ error, actErr error) {
				assert.NoError(t, actErr)
			},
		},
		{
			name: "Negative UserID",
			params: models.TransactionsHistoryParams{
				UserID:   -1,
				SortDate: 1,
				SortSum:  1,
				Page:     1,
				Limit:    1,
			},
			expErr:    errors.New("неверное значение user_id"),
			expReport: "",
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
		{
			name: "Wrong SortDate",
			params: models.TransactionsHistoryParams{
				UserID:   1,
				SortDate: 2,
				SortSum:  1,
				Page:     1,
				Limit:    1,
			},
			expErr:    errors.New("неверное значение sortDate (sortDate=1 - включение сортировки по дате, sortDate=0 - выключение сортировки по дате)"),
			expReport: "",
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
		{
			name: "Wrong SortSum",
			params: models.TransactionsHistoryParams{
				UserID:   1,
				SortDate: 1,
				SortSum:  2,
				Page:     1,
				Limit:    1,
			},
			expErr:    errors.New("неверное значение sortSum (sortSum=1 - включение сортировки по сумме, sortSum=0 - выключение сортировки по сумме)"),
			expReport: "",
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
		{
			name: "Wrong page",
			params: models.TransactionsHistoryParams{
				UserID:   1,
				SortDate: 1,
				SortSum:  1,
				Page:     -1,
				Limit:    1,
			},
			expErr:    errors.New("неверное значение page"),
			expReport: "",
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
		{
			name: "Wrong limit",
			params: models.TransactionsHistoryParams{
				UserID:   1,
				SortDate: 1,
				SortSum:  1,
				Page:     1,
				Limit:    -1,
			},
			expErr:    errors.New("неверное значение limit"),
			expReport: "",
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
		{
			name: "Transactions history returns err",
			params: models.TransactionsHistoryParams{
				UserID:   1,
				SortDate: 1,
				SortSum:  1,
				Page:     1,
				Limit:    1,
			},
			expErr: errors.New("testErr"),
			expDBHistory: []models.TransactionsHistory{
				{},
			},
			expReport: "",
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
		{
			name: "Marshaling returns err",
			params: models.TransactionsHistoryParams{
				UserID:   1,
				SortDate: 1,
				SortSum:  1,
				Page:     1,
				Limit:    1,
			},
			expErr: errors.New("testErr"),
			expDBHistory: []models.TransactionsHistory{
				{Operation: "testOperation",
					Comments: "testComment",
					Time:     time.Unix(0, 0),
					Sum:      1,
				},
			},
			expReport: "",
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockDB := initDatabaseMock(t)
			mockMarshaling := initMarshalCSVMock(t)
			r := NewReport(mockDB, mockMarshaling)

			ctx := context.Background()
			if test.name == "Happy Path" || test.name == "Marshaling returns err" {
				mockDB.EXPECT().TransactionsHistory(ctx, test.params).Return(test.expDBHistory, nil)
				mockMarshaling.EXPECT().MarshalStructHistoryToCSV(test.expDBHistory).Return(test.expReport, test.expErr)
			}
			if test.name == "Transactions history returns err" {
				mockDB.EXPECT().TransactionsHistory(ctx, test.params).Return(test.expDBHistory, test.expErr)
			}

			actReport, actErr := r.TransactionsHistory(ctx, test.params)

			test.assertErr(t, test.expErr, actErr)
			assert.Equal(t, test.expReport, actReport)
		})
	}
}
