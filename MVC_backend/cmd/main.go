package main

import (
    "log"
   "github.com/kartikgoyal137/MVC/pkg/api"
   "github.com/kartikgoyal137/MVC/pkg/models"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    database.SetupDatabase()

    server := api.NewAPIServer(":8080", database.Db)
    if err:=server.Run(); err!=nil {
        log.Fatal(err)
    }
    
}