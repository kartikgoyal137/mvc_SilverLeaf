package workers

import (
	"database/sql"
	"github.com/kartikgoyal137/MVC/pkg/models"
	"time"
)

func StartCacheWorker(db *sql.DB) {
	go func() {
		models.ReloadCategoriesCache(db)
		models.ReloadMenuCache(db)
	}()

	ticker := time.NewTicker(5 * time.Minute)

	go func() {
		for {
			<-ticker.C
			models.ReloadCategoriesCache(db)
			models.ReloadMenuCache(db)
		}
	}()
}