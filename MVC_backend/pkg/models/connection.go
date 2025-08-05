package main

import (
	"database/sql"
	"os"
	"github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {

	godotenv.Load()

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
    cfg.Passwd = os.Getenv("DBPASS")
    cfg.Net = "tcp"
    cfg.Addr = "127.0.0.1:3306"
    cfg.DBName = "velvet_plate"

	var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

	pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}

