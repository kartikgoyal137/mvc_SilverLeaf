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

func (s *CartDB) GetCartItems(orderID int) ([]types.CartItemCheckout, error) {
	rows, err := s.db.Query("SELECT s.order_id, s.product_id, s.quantity, m.price, m.product_name FROM serve AS s JOIN menu AS m ON s.product_id = m.product_id WHERE s.order_id = ?;", orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cart []types.CartItemCheckout
	for rows.Next() {
		c, err := scanRowIntoCartCheckout(rows)
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

func scanRowIntoCartCheckout(rows *sql.Rows) (*types.CartItemCheckout, error) {
	item := new(types.CartItemCheckout)

	err := rows.Scan(
		&item.OrderID,
		&item.ProductID,
		&item.Quantity,
		&item.Price,
		&item.Name,
	)

	if err != nil {
		return nil, err
	}

	return item, nil
}
