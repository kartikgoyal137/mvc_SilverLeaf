package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kartikgoyal137/MVC/pkg/api"
	"github.com/kartikgoyal137/MVC/pkg/models"
	"github.com/kartikgoyal137/MVC/pkg/workers"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	db, err := models.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	workers.InitialCacheLoad(db)
	workers.StartCacheWorker(db)

	server := api.NewAPIServer(":8080", db)

	go func() {
		if err := server.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	if err := models.CloseDatabase(); err != nil {
		log.Printf("Error closing database: %v", err)
	}

	fmt.Println("Server exited gracefully")
}
