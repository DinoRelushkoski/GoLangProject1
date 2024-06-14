package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var database *sql.DB

// GetDB returns the initialized database connection.
func GetDB() *sql.DB {
	return database
}

// Init initializes the SQLite database connection and ensures the 'books' table exists.
func Init() {
	var err error
	database, err = sql.Open("sqlite", "local.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	fmt.Println("Connected to SQLite database successfully!")

	// Set maximum open connections to 10
	database.SetMaxOpenConns(10)

	// Ensure 'books' table exists
	err = createBooksTable()
	if err != nil {
		log.Fatalf("Error creating 'books' table: %v", err)
	}
	fmt.Println("Initialized 'books' table successfully.")
}

// createBooksTable creates the 'books' table if it does not exist.
func createBooksTable() error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		isbn TEXT,
		author TEXT,
		year INTEGER
	)`

	_, err := database.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("error creating 'books' table: %v", err)
	}
	return nil
}
