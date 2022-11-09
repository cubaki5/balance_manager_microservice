package adapter

import (
	"context"
	"errors"
	"time"

	"balance_avito/internal/infrastructure/sql_database/generated/sqlc"
	"balance_avito/internal/models"
	"balance_avito/internal/models/types"
)

func (d *DatabaseAdapter) Accrue(ctx context.Context, accrual models.Accrual) error {
	err := d.execTx(ctx, func(q *sqlc.Queries) error {
		err := q.CreateAccountOrUpdateBalance(ctx, sqlc.CreateAccountOrUpdateBalanceParams{
			UserID:    accrual.UserID.Int64(),
			Balance:   accrual.Income.Int32(),
			Balance_2: accrual.Income.Int32(),
		})

		return err
	})
	return err
}

func (d *DatabaseAdapter) Reservation(ctx context.Context, reservation models.Reservation) error {
	err := d.execTx(ctx, func(q *sqlc.Queries) error {
		dbAccount, err := q.GetAccount(ctx, reservation.UserID.Int64())
		if err != nil {
			return err
		}

		if dbAccount.Balance > reservation.Cost.Int32() {
			err = q.WriteOffMoney(ctx, sqlc.WriteOffMoneyParams{
				Balance: reservation.Cost.Int32(),
				UserID:  reservation.UserID.Int64(),
			})
			if err != nil {
				return err
			}
		} else {
			return errors.New("not enough funds on balance")
		}

		err = q.CreateReservedAccount(ctx, sqlc.CreateReservedAccountParams{
			UserID:    reservation.UserID.Int64(),
			OrderID:   reservation.OrderID.Int64(),
			ServiceID: reservation.ServiceID.Int64(),
			Cost:      reservation.Cost.Int32(),
		})

		return err
	})

	return err
}

func (d *DatabaseAdapter) AcceptPayment(ctx context.Context, reservation models.Reservation) error {
	err := d.execTx(ctx, func(q *sqlc.Queries) error {
		_, err := q.GetReservedAccount(ctx, reservation.OrderID.Int64())
		if err != nil {
			return err
		}

		err = q.DeleteReservedAccount(ctx, reservation.OrderID.Int64())
		if err != nil {
			return err
		}

		err = q.CreateOrUpdateReportRow(ctx, sqlc.CreateOrUpdateReportRowParams{
			ServiceID: reservation.ServiceID.Int64(),
			Year:      int32(time.Now().Year()),
			Month:     int32(time.Now().Month()),
			Income:    reservation.Cost.Int32(),
			Income_2:  reservation.Cost.Int32(),
		})

		return err
	})
	return err
}

func (d *DatabaseAdapter) RejectPayment(ctx context.Context, reservation models.Reservation) error {
	err := d.execTx(ctx, func(q *sqlc.Queries) error {
		_, err := q.GetReservedAccount(ctx, reservation.OrderID.Int64())
		if err != nil {
			return err
		}

		err = q.DeleteReservedAccount(ctx, reservation.OrderID.Int64())
		if err != nil {
			return err
		}

		err = q.ReturnMoney(ctx, sqlc.ReturnMoneyParams{
			Balance: reservation.Cost.Int32(),
			UserID:  reservation.UserID.Int64(),
		})

		return err
	})
	return err
}

func (d *DatabaseAdapter) GetBalance(ctx context.Context, account models.Account) (models.Account, error) {
	err := d.execTx(ctx, func(q *sqlc.Queries) error {
		dbAccount, err := q.GetAccount(ctx, account.UserID.Int64())
		if err != nil {
			account = models.Account{}
			return err
		}

		account.Balance = types.Balance(dbAccount.Balance)
		return err
	})
	return account, err
}
