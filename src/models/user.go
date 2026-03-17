package models

import (
	"database/sql"
	"errors"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (u *User) GetUser(id int, db *sql.DB) (*User, error) {
	// Query the database
	return &User{ID: id, Username: "johnDoe", Email: "john@example.com"}, nil
}
