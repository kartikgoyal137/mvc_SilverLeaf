package models

import (
	"database/sql"
	"github.com/kartikgoyal137/MVC/pkg/types"
)

type MenuDB struct {
	db *sql.DB
}

func NewMenuDB(db *sql.DB) *MenuDB {
	return &MenuDB{db: db}
}

func (s *MenuDB) ListOfCategory() ([]types.Category ,error) {
	rows, err := s.db.Query("SELECT * from categories;")
	if err!=nil {
		return nil, err
	}
	defer rows.Close()

	var cat []types.Category

	for rows.Next() {
		c, err := scanRowIntoCategory(rows)
		if err!=nil {
			return nil, err
		}
		cat = append(cat, *c)
	}

	return cat, nil
}

func (s *MenuDB) GetMenuByCategoryId(id int) ([]types.MenuItem ,error) {
	rows, err := s.db.Query("SELECT * from menu WHERE category_id =  ?;", id)
	if err!=nil {
		return nil, err
	}
	defer rows.Close()

	var items []types.MenuItem

	for rows.Next() {
		i, err := scanRowIntoItem(rows)
		if err!=nil {
			return nil, err
		}
		items = append(items, *i)
	}

	return items , nil
}


func scanRowIntoItem(rows *sql.Rows) (*types.MenuItem, error) {
	item := new(types.MenuItem)

	err := rows.Scan(
		&item.ProductID,
		&item.ProductName,
		&item.CategoryID,
		&item.Price,
		&item.ImageURL,
	)

	if err!=nil {
		return nil,err
	}

	return item, nil
}

func scanRowIntoCategory(rows *sql.Rows) (*types.Category, error) {
	item := new(types.Category)

	err := rows.Scan(
		&item.CategoryID,
		&item.CategoryName,
		&item.ImageURL,
		&item.Description,
	)

	if err!=nil {
		return nil,err
	}

	return item, nil
}


