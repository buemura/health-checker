package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func Connect() {
	DB, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer DB.Close()
}
