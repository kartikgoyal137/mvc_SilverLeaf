package models

import (
	"database/sql"
	"github.com/kartikgoyal137/MVC/pkg/types"
)

type OrderDB struct {
	db *sql.DB
}

func NewOrderDB(db *sql.DB) *OrderDB {
	return &OrderDB{db : db}
}

func (s *OrderDB) GetAllOrders() ([]types.Order, error) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err!=nil {
		return nil, err
	}

	var order []types.Order

	for rows.Next() {
		u, err := scanRowIntoOrder(rows)
		if err!=nil {
			return nil, err
		}
		order = append(order, *u)
	}

	return order, nil
}

func (s *OrderDB) OrdersByStatus(status string) ([]types.Order ,error) {
	rows, err := s.db.Query("SELECT * FROM orders WHERE status = ?", status)
	if err!=nil {
		return nil, err
	}

	var ord []types.Order

	for rows.Next() {
		o, err := scanRowIntoOrder(rows)
		if err!=nil {
			return nil, err
		}
		ord = append(ord, *o)
	}

	

	return ord, nil
}


func (s *OrderDB) OrdersByUserId(id int) ([]types.Order, error) {

	rows, err := s.db.Query("SELECT * FROM orders WHERE user_id = ?", id)
	if err!=nil {
		return nil, err
	}

	var ord []types.Order

	for rows.Next() {
		o, err := scanRowIntoOrder(rows)
		if err!=nil {
			return nil, err
		}
		ord = append(ord, *o)
	}

	

	return ord, nil
}

func (s *OrderDB) UpdateOrder(order types.CreateOrder) error {
	_ , err := s.db.Query("UPDATE orders SET status = ?, instructions = ?, table_no = ? WHERE order_id = ?;", "Yet to Start", order.Instructions, order.TableNo, order.OrderId)
	if err!=nil {
		return err
	}
	return nil
}

func (s *OrderDB) CreateEmptyOrder(user types.User) error {
	_ , err := s.db.Query("INSERT INTO orders (user_id) VALUES (?);", user.UserID)
	if err!=nil {
		return err
	}
	return nil
}

func scanRowIntoOrder(rows *sql.Rows) (*types.Order, error) {
	order := new(types.Order)

	err := rows.Scan(
		&order.OrderID,
		&order.UserID,
		&order.Status,
		&order.CreatedAt,
		&order.Instructions,
		&order.TableNo,
	)

	if err!=nil {
		return nil,err
	}

	return order, nil
}