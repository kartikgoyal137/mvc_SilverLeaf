package models

import (
	"database/sql"
	"github.com/kartikgoyal137/MVC/pkg/types"
)

type PaymentDB struct {
	db *sql.DB
}

func NewPaymentDB(db *sql.DB) *PaymentDB {
	return &PaymentDB{db: db}
}


func (s *PaymentDB) PaymentsByUserId(id int) ([]types.Payment, error) {
	rows, err := s.db.Query("SELECT * FROM payments WHERE user_id = ?", id)
	if err!=nil {
		return nil, err
	}
	defer rows.Close()

	var pay []types.Payment

	for rows.Next() {
		p, err := scanRowIntoPayment(rows)
		if err!=nil {
			return nil, err
		}
		pay = append(pay, *p)

	}

	return pay, nil
}

func (s *PaymentDB) GetAllPayments() ([]types.Payment, error) {
	rows, err := s.db.Query("SELECT * FROM payments;")
	if err!=nil {
		return nil, err
	}
	defer rows.Close()

	var pay []types.Payment

	for rows.Next() {
		p, err := scanRowIntoPayment(rows)
		if err!=nil {
			return nil, err
		}
		pay = append(pay, *p)

	}

	return pay, nil
}

func (s *PaymentDB) CreateNewPayment(pay *types.Payment) error {
	_ , err := s.db.Query("INSERT INTO payments (order_id, user_id, food_total, created_at, tip) VALUES (?, ?, ?, ?, ?);", pay.OrderID, pay.UserID, pay.FoodTotal, pay.CreatedAt, pay.Tip)

	if err!=nil {
		return err
	}

	return nil
}


func scanRowIntoPayment(rows *sql.Rows) (*types.Payment, error) {
	pay := new(types.Payment)

	err := rows.Scan(
		&pay.TransactionID,
		&pay.OrderID,
		&pay.UserID,
		&pay.FoodTotal,
		&pay.CreatedAt,
		&pay.Tip,
		&pay.Status,
	)

	if err!=nil {
		return nil,err
	}

	return pay, nil
}