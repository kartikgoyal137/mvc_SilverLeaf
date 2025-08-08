package models

import (
	"database/sql"
	"fmt"
	"github.com/kartikgoyal137/MVC/pkg/types"
)

const userColumns = "user_id, first_name, last_name, contact, email, password_hash, role"

type UserDB struct {
	db *sql.DB
}

func NewUserDB(db *sql.DB) *UserDB {
	return &UserDB{db: db}
}

func (s *UserDB) GetAllUsers() ([]types.User, error) {
	query := fmt.Sprintf("SELECT %s FROM users", userColumns)
	rows, err := s.db.Query(query)
	if err!=nil {
		return nil, err
	}
	defer rows.Close()

	var people []types.User

	for rows.Next() {
		u, err := scanRowIntoUser(rows)
		if err!=nil {
			return nil, err
		}
		people = append(people, *u)
	}

	return people, nil
}

func (s *UserDB) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err!=nil {
		return nil, err
	}
	defer rows.Close()

	var u *types.User

	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err!=nil {
			return nil, err
		}
	}

	if u == nil {
        return nil, fmt.Errorf("user not found")
    }

	return u, nil
}

func (s *UserDB) CreateNewUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (first_name, last_name, contact, email, password_hash) VALUES (?, ?, ?, ?, ?)", user.FirstName, user.LastName, user.Contact ,user.Email, user.PasswordHash)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserDB) GetUserById(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE user_id = ?", id)
	if err!=nil {
		return nil, err
	}
	defer rows.Close()

	u := new(types.User)

	for rows.Next() {
		u, err= scanRowIntoUser(rows)
		if err!=nil {
			return nil, err
		}
	}

	if u.UserID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}


func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.UserID,
		&user.FirstName,
		&user.LastName,
		&user.Contact,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
	)

	if err!=nil {
		return nil,err
	}

	return user, nil
}

