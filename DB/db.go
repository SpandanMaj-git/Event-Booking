package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err := sql.Open("sqlite", "api.db")

	if err != nil {
		panic("failed to connect database")
	}

	DB.SetConnMaxIdleTime(5)
	DB.SetMaxOpenConns(10)

	createTables()
}

func createTables() {
	createEventsTables := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER	
	)
    `

	_, err := DB.Exec(createEventsTables)

	if err != nil {
		panic("failed to create events table")
	}
}
