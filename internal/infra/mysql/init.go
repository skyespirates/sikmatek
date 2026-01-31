package mysql

import (
	"context"
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
