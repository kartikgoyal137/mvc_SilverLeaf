package workers

import (
	"database/sql"
	"github.com/kartikgoyal137/MVC/pkg/models"
	"time"
	"log"
)

func InitialCacheLoad(db *sql.DB) {
	log.Println("Performing initial synchronous cache load...")
	models.ReloadCategoriesCache(db)
	models.ReloadMenuCache(db)
	log.Println("Initial cache load complete. Cache is ready.")
}

func StartCacheWorker(db *sql.DB) {

	ticker := time.NewTicker(5 * time.Minute)

	go func() {
		for {
			<-ticker.C
			models.ReloadCategoriesCache(db)
			models.ReloadMenuCache(db)
		}
	}()
}