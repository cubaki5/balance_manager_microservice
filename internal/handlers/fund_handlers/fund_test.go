package fund_handlers

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"balance_avito/internal/handlers/fund_handlers/mocks"
	"balance_avito/internal/models"
	"balance_avito/internal/models/types"
)

const (
	happyPath          = "HappyPath"
	interfaceReturnErr = "Interface returns err"
)

func initDatabaseMock(t *testing.T) *mocks.MockDatabase {
	ctrl := gomock.NewController(t)
	return mocks.NewMockDatabase(ctrl)
}

func TestFund_Accrue(t *testing.T) {
	tests := []struct {
		name      string
		accrual   models.Accrual
		expErr    error
		assertErr func(t *testing.T, expErr error, actErr error)
	}{
		{
			name: happyPath,
			accrual: models.Accrual{
				UserID: 1,
				Income: 1,
			},
			expErr: nil,
			assertErr: func(t *testing.T, _ error, actErr error) {
				assert.NoError(t, actErr)
			},
		},
		{
			name: "Negative UserID",
			accrual: models.Accrual{
				UserID: -1,
				Income: 1,
			},
			expErr: errors.New("неверное значение user_id"),
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
		{
			name: "Negative Income",
			accrual: models.Accrual{
				UserID: 1,
				Income: -1,
			},
			expErr: errors.New("неверное значение income"),
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
		{
			name: "Accrue returns err",
			accrual: models.Accrual{
				UserID: 1,
				Income: 1,
			},
			expErr: errors.New("testErr"),
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockDB := initDatabaseMock(t)
			fund := NewFund(mockDB)

			ctx := context.Background()
			if test.name == happyPath || test.name == "Accrue returns err" {
				mockDB.EXPECT().Accrue(ctx, test.accrual).Return(test.expErr)
			}
			actErr := fund.Accrue(ctx, test.accrual)

			test.assertErr(t, test.expErr, actErr)
		})
	}
}

func TestFund_Reservation_Accept_RejectPayment(t *testing.T) {
	tests := []struct {
		name        string
		reservation models.Reservation
		expErr      error
		assertErr   func(t *testing.T, expErr error, actErr error)
	}{
		{
			name: happyPath,
			reservation: models.Reservation{
				UserID:    1,
				ServiceID: 1,
				OrderID:   1,
				Cost:      1,
			},
			expErr: nil,
			assertErr: func(t *testing.T, _ error, actErr error) {
				assert.NoError(t, actErr)
			},
		},
		{
			name: "Negative UserID",
			reservation: models.Reservation{
				UserID:    -1,
				ServiceID: 1,
				OrderID:   1,
				Cost:      1,
			},
			expErr: errors.New("неверное значение user_id"),
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
		{
			name: "Negative ServiceID",
			reservation: models.Reservation{
				UserID:    1,
				ServiceID: -1,
				OrderID:   1,
				Cost:      1,
			},
			expErr: errors.New("неверное значение service_id"),
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
		{
			name: "Negative OrderID",
			reservation: models.Reservation{
				UserID:    1,
				ServiceID: 1,
				OrderID:   -1,
				Cost:      1,
			},
			expErr: errors.New("неверное значение order_id"),
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
		{
			name: "Negative cost",
			reservation: models.Reservation{
				UserID:    1,
				ServiceID: 1,
				OrderID:   1,
				Cost:      -1,
			},
			expErr: errors.New("неверное значение cost"),
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
		{
			name: interfaceReturnErr,
			reservation: models.Reservation{
				UserID:    1,
				ServiceID: 1,
				OrderID:   1,
				Cost:      1,
			},
			expErr: errors.New("testErr"),
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockDB := initDatabaseMock(t)
			fund := NewFund(mockDB)

			ctx := context.Background()
			if test.name == happyPath || test.name == interfaceReturnErr {
				mockDB.EXPECT().Reservation(ctx, test.reservation).Return(test.expErr)
			}
			actErr := fund.Reservation(ctx, test.reservation)

			test.assertErr(t, test.expErr, actErr)
		})
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockDB := initDatabaseMock(t)
			fund := NewFund(mockDB)

			ctx := context.Background()
			if test.name == happyPath || test.name == interfaceReturnErr {
				mockDB.EXPECT().RejectPayment(ctx, test.reservation).Return(test.expErr)
			}
			actErr := fund.RejectPayment(ctx, test.reservation)

			test.assertErr(t, test.expErr, actErr)
		})
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockDB := initDatabaseMock(t)
			fund := NewFund(mockDB)

			ctx := context.Background()
			if test.name == "Happy Path" || test.name == "Interface returns err" {
				mockDB.EXPECT().AcceptPayment(ctx, test.reservation).Return(test.expErr)
			}
			actErr := fund.AcceptPayment(ctx, test.reservation)

			test.assertErr(t, test.expErr, actErr)
		})
	}
}

func TestFund_GetBalance(t *testing.T) {
	tests := []struct {
		name       string
		account    models.Account
		expErr     error
		expBalance types.Balance
		assertErr  func(t *testing.T, expErr error, actErr error)
	}{
		{
			name: happyPath,
			account: models.Account{
				UserID: 1,
			},
			expErr:     nil,
			expBalance: 1,
			assertErr: func(t *testing.T, _ error, actErr error) {
				assert.NoError(t, actErr)
			},
		},
		{
			name: "Negative UserID",
			account: models.Account{
				UserID: -1,
			},
			expErr:     errors.New("неверное значение user_id"),
			expBalance: 0,
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
		{
			name: "GetBalance returns err",
			account: models.Account{
				UserID: 1,
			},
			expErr:     errors.New("testErr"),
			expBalance: 0,
			assertErr: func(t *testing.T, expErr error, actErr error) {
				assert.EqualError(t, actErr, expErr.Error())
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockDB := initDatabaseMock(t)
			fund := NewFund(mockDB)

			ctx := context.Background()
			if test.name == happyPath || test.name == "GetBalance returns err" {
				mockDB.EXPECT().GetBalance(ctx, test.account).Return(models.Account{
					UserID:  test.account.UserID,
					Balance: 1,
				}, test.expErr)
			}
			actBalance, actErr := fund.GetBalance(ctx, test.account)

			test.assertErr(t, test.expErr, actErr)
			assert.Equal(t, test.expBalance, actBalance)
		})
	}
}
