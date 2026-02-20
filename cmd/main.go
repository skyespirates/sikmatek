package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
	"github.com/skyespirates/sikmatek/internal/infra/mysql"
	"github.com/skyespirates/sikmatek/internal/logger"
)

type application struct {
	logger *logger.Logger
	db     *sql.DB
	c      *cloudinary.Cloudinary
}

func main() {
	godotenv.Load()

	cld, err := cloudinary.New()
	if err != nil {
		log.Fatal("connection failed to cloudinary")
	}
	log.Println("connected to cloudinary")

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
		c:      cld,
	}

	app.serve()
}
