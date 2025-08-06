package models

import (
	"database/sql"
	"fmt"
	"github.com/kartikgoyal137/MVC/pkg/types"
)

type OrderStore struct {
	db *sql.DB
}

func (s *OrderStore) NewStore(db *sql.DB) *OrderStore {
	return &OrderStore{db : db}
}

func (s *OrderStore) OrdersByStatus() {

}

func (s *OrderStore) OrderIDinServe() {

}

func (s *OrderStore) OrdersByUserId() {

}

func (s *OrderStore) CreateOrder() {
	rows, err = s.db.Query()
}