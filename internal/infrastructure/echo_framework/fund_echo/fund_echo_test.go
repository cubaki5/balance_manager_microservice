package fund_echo

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"balance_avito/internal/infrastructure/echo_framework/fund_echo/mocks"
	"balance_avito/internal/models"
	"balance_avito/internal/models/types"
)

func initMockFund(t *testing.T) *mocks.MockFund {
	ctr := gomock.NewController(t)
	return mocks.NewMockFund(ctr)
}

const bindRetErr = "BindReturnsErr"

func TestHandler_Accrue(t *testing.T) {
	tests := []struct {
		name       string
		json       string
		accrual    models.Accrual
		expErr     error
		expRetJSON string
		statusCode int
	}{
		{
			name: "HappyPath",
			json: `{"user_id":3,"income":1500}`,
			accrual: models.Accrual{
				UserID: 3,
				Income: 1500,
			},
			expErr:     nil,
			expRetJSON: "",
			statusCode: http.StatusOK,
		},
		{
			name: "AccrueReturnsErr",
			json: `{"user_id":3,"income":1500}`,
			accrual: models.Accrual{
				UserID: 3,
				Income: 1500,
			},
			expErr:     errors.New("testErr"),
			expRetJSON: "\"testErr\"\n",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       bindRetErr,
			json:       `wrongJSON`,
			expRetJSON: "\"code=400, message=Syntax error: offset=1, error=invalid character 'w' looking for beginning of value, internal=invalid character 'w' looking for beginning of value\"\n",
			statusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockFund := initMockFund(t)
			fundEcho := NewFundHandler(mockFund)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/fund/accrual", strings.NewReader(test.json))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if test.name != bindRetErr {
				mockFund.EXPECT().Accrue(c.Request().Context(), test.accrual).Return(test.expErr)
			}

			actErr := fundEcho.Accrue(c)

			assert.NoError(t, actErr)
			assert.Equal(t, test.statusCode, rec.Code)
			assert.Equal(t, test.expRetJSON, rec.Body.String())
		})
	}
}

func TestHandler_Reservation_Accept_RejectPayment(t *testing.T) {
	tests := []struct {
		name        string
		json        string
		reservation models.Reservation
		expErr      error
		expRetJSON  string
		statusCode  int
	}{
		{
			name: "HappyPath",
			json: `{
				  "user_id":3,
				  "service_id":666,
				  "order_id":83,
				  "cost":300
			}`,
			reservation: models.Reservation{
				UserID:    3,
				ServiceID: 666,
				OrderID:   83,
				Cost:      300,
			},
			expErr:     nil,
			expRetJSON: "",
			statusCode: http.StatusOK,
		},
		{
			name: "ReturnsErr",
			json: `{
				  "user_id":3,
				  "service_id":666,
				  "order_id":83,
				  "cost":300
			}`,
			reservation: models.Reservation{
				UserID:    3,
				ServiceID: 666,
				OrderID:   83,
				Cost:      300,
			},
			expErr:     errors.New("testErr"),
			expRetJSON: "\"testErr\"\n",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       bindRetErr,
			json:       `wrongJSON`,
			expRetJSON: "\"code=400, message=Syntax error: offset=1, error=invalid character 'w' looking for beginning of value, internal=invalid character 'w' looking for beginning of value\"\n",
			statusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run("Reservation"+test.name, func(t *testing.T) {

			mockFund := initMockFund(t)
			fundEcho := NewFundHandler(mockFund)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/fund/reservation", strings.NewReader(test.json))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if test.name != bindRetErr {
				mockFund.EXPECT().Reservation(c.Request().Context(), test.reservation).Return(test.expErr)
			}

			actErr := fundEcho.Reservation(c)

			assert.NoError(t, actErr)
			assert.Equal(t, test.statusCode, rec.Code)
			assert.Equal(t, test.expRetJSON, rec.Body.String())
		})
	}
	for _, test := range tests {
		t.Run("AcceptPay"+test.name, func(t *testing.T) {

			mockFund := initMockFund(t)
			fundEcho := NewFundHandler(mockFund)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/fund/payment_acceptance", strings.NewReader(test.json))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if test.name != bindRetErr {
				mockFund.EXPECT().AcceptPayment(c.Request().Context(), test.reservation).Return(test.expErr)
			}

			actErr := fundEcho.AcceptPayment(c)

			assert.NoError(t, actErr)
			assert.Equal(t, test.statusCode, rec.Code)
			assert.Equal(t, test.expRetJSON, rec.Body.String())
		})
	}
	for _, test := range tests {
		t.Run("RejectPay"+test.name, func(t *testing.T) {

			mockFund := initMockFund(t)
			fundEcho := NewFundHandler(mockFund)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/fund/payment_rejection", strings.NewReader(test.json))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if test.name != bindRetErr {
				mockFund.EXPECT().RejectPayment(c.Request().Context(), test.reservation).Return(test.expErr)
			}

			actErr := fundEcho.RejectPayment(c)

			assert.NoError(t, actErr)
			assert.Equal(t, test.statusCode, rec.Code)
			assert.Equal(t, test.expRetJSON, rec.Body.String())
		})
	}
}

func TestHandler_GetBalance(t *testing.T) {
	tests := []struct {
		name           string
		json           string
		account        models.Account
		expErr         error
		expBalance     types.Balance
		expJSONBalance string
		statusCode     int
	}{
		{
			name:           "HappyPath",
			json:           `{"user_id":3}`,
			account:        models.Account{UserID: 3},
			expErr:         nil,
			expBalance:     1500,
			expJSONBalance: "1500\n",
			statusCode:     http.StatusOK,
		},
		{
			name:           "AccrueReturnsErr",
			json:           `{"user_id":3}`,
			account:        models.Account{UserID: 3},
			expErr:         errors.New("testErr"),
			expBalance:     0,
			expJSONBalance: "\"testErr\"\n",
			statusCode:     http.StatusBadRequest,
		},
		{
			name:           bindRetErr,
			json:           `wrongJSON`,
			expJSONBalance: "\"code=400, message=Syntax error: offset=1, error=invalid character 'w' looking for beginning of value, internal=invalid character 'w' looking for beginning of value\"\n",
			statusCode:     http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockFund := initMockFund(t)
			fundEcho := NewFundHandler(mockFund)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/fund/balance", strings.NewReader(test.json))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if test.name != bindRetErr {
				mockFund.EXPECT().GetBalance(c.Request().Context(), test.account).Return(test.expBalance, test.expErr)
			}

			actErr := fundEcho.GetBalance(c)

			assert.NoError(t, actErr)
			assert.Equal(t, test.statusCode, rec.Code)
			assert.Equal(t, test.expJSONBalance, rec.Body.String())
		})
	}
}
