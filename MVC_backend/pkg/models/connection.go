package database

import (
	"database/sql"
	"os"
	"github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"github.com/joho/godotenv"
)

var Db *sql.DB

func SetupDatabase() {
    
    
	godotenv.Load()

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
    cfg.Passwd = os.Getenv("DBPASS")
    cfg.Net = "tcp"
    cfg.Addr = "127.0.0.1:3306"
    cfg.DBName = "velvet_plate"

	var err error
    Db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

	pingErr := Db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}

