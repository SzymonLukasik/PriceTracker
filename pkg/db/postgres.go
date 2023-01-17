package db

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func ExecuteTransaction(ctx context.Context, db *sql.DB, f func() error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.WithError(err).Error("unable to begin transaction")
		return err
	}
	err = f()
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}
