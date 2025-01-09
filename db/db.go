package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const dbPath = "./user.db"

func init() {
	// Check if DB exists.
	_, err := os.Stat(dbPath)
	if err == nil {
		fmt.Println("DB Exists, exiting.")
		setDB()
		return
	}
	fmt.Println("DB does not exist, creating.")
	// Set up connection.
	setDB()
	// Execute schema and insert starting data.
	err = initDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("DB Created successfully.")

}

func setDB() {
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	// Engage WAL mode for MAX PERFORMANCE
	_, err = db.Exec("PRAGMA journal_mode = WAL;")
	if err != nil {
		log.Fatalf("Failed to set WAL mode: %v", err)
	}
}

func GetDB() *sql.DB {
	return db
}

func initDB() error {
	schema, err := os.ReadFile("./schema.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	schemaString := string(schema)
	_, err = db.Exec(schemaString)
	if err != nil {
		return err
	}
	// No error if nil.
	return nil
}
