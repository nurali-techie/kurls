package utils

import (
	"database/sql"
	"fmt"
	"os"
)

func GetDatabase() *sql.DB {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dbFile := fmt.Sprintf("%s/.kurls/kurls.db", homeDir)
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		panic(err)
	}
	return db
}

func InitDB(db *sql.DB) error {
	if db == nil {
		return fmt.Errorf("db not available")
	}
	err := createTable(db, "CREATE TABLE IF NOT EXISTS kurls (key TEXT PRIMARY KEY, curl TEXT)")
	if err != nil {
		return err
	}
	err = createTable(db, "CREATE TABLE IF NOT EXISTS memory (seq INTEGER PRIMARY KEY, key TEXT)")
	return err
}

func createTable(db *sql.DB, query string) error {
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}
