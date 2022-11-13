package report_echo

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"balance_avito/internal/infrastructure/echo_framework/report_echo/mocks"
	"balance_avito/internal/models"
)

func initMockReport(t *testing.T) *mocks.MockReport {
	ctr := gomock.NewController(t)
	return mocks.NewMockReport(ctr)
}

func TestHandler_Accounting(t *testing.T) {
	tests := []struct {
		name          string
		json          string
		reportDate    models.ReportDate
		expErr        error
		expReport     string
		expJSONReport string
		statusCode    int
	}{
		{
			name: "HappyPath",
			json: `{"year":2020,"month":11}`,
			reportDate: models.ReportDate{
				Year:  2020,
				Month: 11,
			},
			expErr:        nil,
			expReport:     "43;555000\n33;554222",
			expJSONReport: "43;555000\n33;554222",
			statusCode:    http.StatusOK,
		},
		{
			name: "AccountingReturnsErr",
			json: `{"year":2020,"month":11}`,
			reportDate: models.ReportDate{
				Year:  2020,
				Month: 11,
			},
			expErr:        errors.New("testErr"),
			expReport:     "",
			expJSONReport: "\"testErr\"\n",
			statusCode:    http.StatusBadRequest,
		},
		{
			name:          "BindReturnsErr",
			json:          `wrongJSON`,
			expJSONReport: "\"code=400, message=Syntax error: offset=1, error=invalid character 'w' looking for beginning of value, internal=invalid character 'w' looking for beginning of value\"\n",
			statusCode:    http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockReport := initMockReport(t)
			fundEcho := NewReportHandler(mockReport)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/report/accounting", strings.NewReader(test.json))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if test.name != "BindReturnsErr" {
				mockReport.EXPECT().Accounting(c.Request().Context(), test.reportDate).Return(test.expReport, test.expErr)
			}

			actErr := fundEcho.Accounting(c)

			assert.NoError(t, actErr)
			assert.Equal(t, test.statusCode, rec.Code)
			assert.Equal(t, test.expJSONReport, rec.Body.String())
		})
	}
}

func TestHandler_TransactionsHistory(t *testing.T) {
	tests := []struct {
		name           string
		json           string
		historyParams  models.TransactionsHistoryParams
		expErr         error
		expHistory     string
		expJSONHistory string
		statusCode     int
	}{
		{
			name: "HappyPath",
			json: `{"user_id":3}`,
			historyParams: models.TransactionsHistoryParams{
				UserID:   3,
				SortDate: 1,
				SortSum:  0,
				Page:     1,
				Limit:    100,
			},
			expErr:         nil,
			expHistory:     "test;test\ntest;test",
			expJSONHistory: "test;test\ntest;test",
			statusCode:     http.StatusOK,
		},
		{
			name: "AccrueReturnsErr",
			json: `{"user_id":3}`,
			historyParams: models.TransactionsHistoryParams{
				UserID:   3,
				SortDate: 1,
				SortSum:  0,
				Page:     1,
				Limit:    100,
			},
			expErr:         errors.New("testErr"),
			expHistory:     "test;test\ntest;test",
			expJSONHistory: "\"testErr\"\n",
			statusCode:     http.StatusBadRequest,
		},
		{
			name:           "BindReturnsErr",
			json:           `wrongJSON`,
			expJSONHistory: "\"code=400, message=Syntax error: offset=1, error=invalid character 'w' looking for beginning of value, internal=invalid character 'w' looking for beginning of value\"\n",
			statusCode:     http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockReport := initMockReport(t)
			fundEcho := NewReportHandler(mockReport)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/report/transactions_history?sortDate=1&sortSum=0&page=1&limit=100", strings.NewReader(test.json))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if test.name != "BindReturnsErr" {
				mockReport.EXPECT().TransactionsHistory(c.Request().Context(), test.historyParams).Return(test.expHistory, test.expErr)

			}

			actErr := fundEcho.TransactionsHistory(c)

			assert.NoError(t, actErr)
			assert.Equal(t, test.statusCode, rec.Code)
			assert.Equal(t, test.expJSONHistory, rec.Body.String())
		})
	}
}
