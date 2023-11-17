package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySQL Connection Driver
)

// Database connection
func Connect() (*sql.DB, error) {

	connectionString := "devbook:devbook@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, error := sql.Open("mysql", connectionString)

	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		return nil, error
	}

	return db, nil
}
