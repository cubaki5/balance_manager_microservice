// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: accounts.sql

package sqlc

import (
	"context"
)

const createAccountOrUpdateBalance = `-- name: CreateAccountOrUpdateBalance :exec
INSERT INTO accounts (
    user_id, balance
) VALUE (
         ?, ?
    )
ON DUPLICATE KEY UPDATE balance = balance + ?
`

type CreateAccountOrUpdateBalanceParams struct {
	UserID    int64
	Balance   int32
	Balance_2 int32
}

func (q *Queries) CreateAccountOrUpdateBalance(ctx context.Context, arg CreateAccountOrUpdateBalanceParams) error {
	_, err := q.db.ExecContext(ctx, createAccountOrUpdateBalance, arg.UserID, arg.Balance, arg.Balance_2)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT user_id, balance FROM accounts
WHERE user_id = ? LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, userID int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, userID)
	var i Account
	err := row.Scan(&i.UserID, &i.Balance)
	return i, err
}

const returnMoney = `-- name: ReturnMoney :exec
UPDATE accounts
SET balance = balance + ?
WHERE user_id = ?
`

type ReturnMoneyParams struct {
	Balance int32
	UserID  int64
}

func (q *Queries) ReturnMoney(ctx context.Context, arg ReturnMoneyParams) error {
	_, err := q.db.ExecContext(ctx, returnMoney, arg.Balance, arg.UserID)
	return err
}

const writeOffMoney = `-- name: WriteOffMoney :exec
UPDATE accounts
SET balance = balance - ?
WHERE user_id = ?
`

type WriteOffMoneyParams struct {
	Balance int32
	UserID  int64
}

func (q *Queries) WriteOffMoney(ctx context.Context, arg WriteOffMoneyParams) error {
	_, err := q.db.ExecContext(ctx, writeOffMoney, arg.Balance, arg.UserID)
	return err
}
