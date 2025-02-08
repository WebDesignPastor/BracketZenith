package database

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./bracketZenith.db")
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to SQLite successfully")
	return db, nil
}
