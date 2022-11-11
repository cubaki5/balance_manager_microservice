package adapter

import (
	"context"
	"database/sql"

	"balance_avito/internal/infrastructure/sql_database/generated/sqlc"

	_ "github.com/go-sql-driver/mysql" // Needed for working with db
	"github.com/labstack/gommon/log"
)

type DatabaseAdapter struct {
	*sqlc.Queries
	db *sql.DB
}

func NewDatabaseAdapter(db *sql.DB) *DatabaseAdapter {
	return &DatabaseAdapter{
		Queries: sqlc.New(db),
		db:      db,
	}
}

func RunDB() (*sql.DB, error) {
	db, err := sql.Open(mysql, dataSourceName)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *DatabaseAdapter) execTx(ctx context.Context, fn func(queries *sqlc.Queries) error) error {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := sqlc.New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Error(rbErr)
		}
		return err
	}

	return tx.Commit()
}
