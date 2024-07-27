package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func Connect() {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PASS, config.DB_DATABASE)
	// DB, err := sql.Open("postgres", psqlInfo)

	DB, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer DB.Close()
}
