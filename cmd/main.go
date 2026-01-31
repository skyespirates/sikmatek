package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/skyespirates/sikmatek/internal/infra/mysql"
	"github.com/skyespirates/sikmatek/internal/logger"
)

type application struct {
	logger *logger.Logger
	db     *sql.DB
}

func main() {
	godotenv.Load()

	db, err := mysql.InitDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	log.Println("database connection pool established")

	logger := logger.New(os.Stdout)

	app := &application{
		logger: logger,
		db:     db,
	}

	app.serve()
}
