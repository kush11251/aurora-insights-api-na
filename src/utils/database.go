package utils

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDatabase() (*sqlx.DB, error) {
	dsn := "user:password@tcp(localhost:3306)/database"
	db, err := sqlx.Connect("mysql", dsn)
	return db, err
}
