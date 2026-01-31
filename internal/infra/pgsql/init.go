package pgsql

import (
	"context"
	"database/sql"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DSN"))
	if err != nil {
		return nil, err
	}

	moc := os.Getenv("MAX_OPEN_CONNS")
	maxOpenConns, _ := strconv.Atoi(moc)
	mic := os.Getenv("MAX_IDLE_CONNS")
	maxIdleConns, _ := strconv.Atoi(mic)

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	duration, err := time.ParseDuration(os.Getenv("MAX_IDLE_TIME"))
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
