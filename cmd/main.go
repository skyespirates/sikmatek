package main

import (
	"database/sql"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
	"github.com/skyespirates/sikmatek/internal/infra/mysql"
	"github.com/skyespirates/sikmatek/internal/utils"
)

type application struct {
	logger *utils.Logger
	db     *sql.DB
	c      *cloudinary.Cloudinary
}

func main() {
	godotenv.Load()
	logger := utils.New(os.Stdout, utils.LevelInfo)

	cld, err := cloudinary.New()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	logger.PrintInfo("connected to cloudinary", nil)

	db, err := mysql.InitDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	logger.PrintInfo("database connection pool established", nil)

	app := &application{
		logger: logger,
		db:     db,
		c:      cld,
	}

	app.serve()
}
