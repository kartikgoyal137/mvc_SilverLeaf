package models

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

var DB *sql.DB

func InitDatabase() (*sql.DB, error) {

	err2 := godotenv.Load()
	if err2!=nil {
		return nil, err2
	}

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.DBName = os.Getenv("DBNAME")
	cfg.ParseTime = true
	dbHost := os.Getenv("DBHOST")
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}
	cfg.Addr = dbHost + ":3306"

	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		return nil, pingErr
	}
	fmt.Println("Connected!")

	return DB, nil
}

func CloseDatabase() error {
	if DB != nil {
		fmt.Println("Closing database connection...")
		return DB.Close()
	}
	return nil
}
