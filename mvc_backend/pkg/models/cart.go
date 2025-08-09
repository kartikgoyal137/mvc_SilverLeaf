package models

import (
	"database/sql"
	"github.com/kartikgoyal137/MVC/pkg/types"
)

type CartDB struct {
	db *sql.DB
}

func NewCartDB(db *sql.DB) *CartDB {
	return &CartDB{db: db}
}

func (s *CartDB) AddToCart(place types.CartItem) error {
	_, err := s.db.Exec("INSERT INTO serve (order_id, product_id, quantity) VALUES (?, ?, ?);", place.OrderID, place.ProductID, place.Quantity)

	if err != nil {
		return err
	}

	return nil

}

func (s *CartDB) GetCartItems(orderID int) ([]types.CartItem, error) {
	rows, err := s.db.Query("SELECT * FROM serve WHERE order_id = ?", orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cart []types.CartItem
	for rows.Next() {
		c, err := scanRowIntoCart(rows)
		if err != nil {
			return nil, err
		}
		cart = append(cart, *c)
	}

	return cart, nil
}

func (s *CartDB) UpdateCartItemQuantity(place types.CartItem) error {
	_, err := s.db.Exec("UPDATE serve SET quantity = ? WHERE order_id = ? AND product_id = ?;", place.Quantity, place.OrderID, place.ProductID)

	if err != nil {
		return err
	}

	return nil
}

func (s *CartDB) DeleteCartItem(place types.CartItem) error {
	_, err := s.db.Exec("DELETE FROM serve WHERE order_id = ? AND product_id = ?;", place.OrderID, place.ProductID)

	if err != nil {
		return err
	}

	return nil
}

func scanRowIntoCart(rows *sql.Rows) (*types.CartItem, error) {
	item := new(types.CartItem)

	err := rows.Scan(
		&item.OrderID,
		&item.ProductID,
		&item.Quantity,
	)

	if err != nil {
		return nil, err
	}

	return item, nil
}
