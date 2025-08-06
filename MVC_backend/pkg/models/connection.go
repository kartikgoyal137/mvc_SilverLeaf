package models

import (
	"database/sql"
	"os"
	"github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDatabase() error {
    
    
	godotenv.Load()

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
    cfg.Passwd = os.Getenv("DBPASS")
    cfg.Net = "tcp"
    cfg.Addr = "127.0.0.1:3306"
    cfg.DBName = "velvet_plate"

	var err error
    DB, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        return err
    }

	pingErr := DB.Ping()
    if pingErr != nil {
        return err
    }
    fmt.Println("Connected!")

    return nil
}

