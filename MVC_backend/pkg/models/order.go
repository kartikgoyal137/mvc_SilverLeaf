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
	rows, err := s.db.Query("SELECT * FROM orders")
	if err!=nil {
		return nil, err
	}
	defer rows.Close()

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
	defer rows.Close()

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
	defer rows.Close()

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

func (s *OrderDB) CreateOrder(order types.CreateOrder) error {
	_ , err := s.db.Exec("UPDATE orders SET instructions = ?, table_no = ? WHERE order_id = ?;",order.Instructions, order.TableNo, order.OrderID)
	if err!=nil {
		return err
	}
	
	return nil
}

func (s *OrderDB) CreateEmptyOrder(userId int) (int,error) {
	result , err := s.db.Exec("INSERT INTO orders (user_id, status, instructions, table_no) VALUES (?, 'Yet to start', NULL, NULL);", userId)
	if err!=nil {
		return 0,err
	}

	id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return int(id), nil
}

func (s *OrderDB) ChangeStatus(orderId int, status string) error {
	_ , err := s.db.Exec("UPDATE orders SET status = ? WHERE order_id = ?;", status, orderId)
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

func scanRowIntoServe(rows *sql.Rows) (*types.CartItem, error) {
	order := new(types.CartItem)

	err := rows.Scan(
		&order.OrderID,
		&order.ProductID,
		&order.Quantity,
	)

	if err!=nil {
		return nil,err
	}

	return order, nil
}