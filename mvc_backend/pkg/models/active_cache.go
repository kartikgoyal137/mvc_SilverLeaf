package models

import (
	"database/sql"
	json "github.com/json-iterator/go"
	"log"
	"sync"
	"github.com/kartikgoyal137/MVC/pkg/types"
)

var CategoryCacheString string
var MenuCache = make(map[int]string)
var CacheMutex sync.RWMutex

func ReloadCategoriesCache(db *sql.DB) {
	rows, err := db.Query("SELECT category_id, category_name, description, image_url FROM categories;")
	if err != nil {
		log.Printf("Failed to reload categories cache: %v", err)
		return
	}
	defer rows.Close()

	var categories []types.Category
	for rows.Next() {
		var c types.Category
		if err := rows.Scan(&c.CategoryID, &c.CategoryName, &c.Description, &c.ImageURL); err != nil {
			log.Printf("Failed to scan category row: %v", err)
			continue
		}
		categories = append(categories, c)
	}

	jsonData, err := json.Marshal(categories)
	if err != nil {
		log.Printf("Failed to marshal categories cache: %v", err)
		return
	}

	CacheMutex.Lock()
	CategoryCacheString = string(jsonData)
	CacheMutex.Unlock()

	log.Println("All categories reloaded.")
}

func ReloadMenuCache(db *sql.DB) {
	query := `SELECT 
                m.product_id, 
                m.product_name, 
                m.category_id, 
                m.price, 
                m.image_url, 
                COALESCE(pi.ingredients, '') as ingredients
              FROM menu AS m
              LEFT JOIN product_ingredients AS pi ON m.product_id = pi.product_id
              ORDER BY m.category_id;`

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Failed to execute menu query: %v", err)
		return
	}
	defer rows.Close()

	menuByCategory := make(map[int][]types.MenuItem)

	for rows.Next() {
		var i types.MenuItem
		if err := rows.Scan(&i.ProductID, &i.ProductName, &i.CategoryID, &i.Price, &i.ImageURL, &i.IngredientList); err != nil {
			log.Printf("Failed to scan menu item row %v", err)
			continue
		}
		menuByCategory[i.CategoryID] = append(menuByCategory[i.CategoryID], i)
	}

	tempMenuCache := make(map[int]string)
	for categoryID, items := range menuByCategory {
		jsonData, err := json.Marshal(items)
		if err != nil {
			log.Printf("Failed to marshal menu for category %d: %v", categoryID, err)
			continue
		}
		tempMenuCache[categoryID] = string(jsonData)
	}

	CacheMutex.Lock()
	MenuCache = tempMenuCache
	CacheMutex.Unlock()

	log.Println("Full menu reloaded")
}
