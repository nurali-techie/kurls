package utils

import (
	"database/sql"
	"fmt"
)

func GetDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./kurls.db")
	if err != nil {
		panic(err)
	}
	return db
}

func InitDB(db *sql.DB) error {
	if db == nil {
		return fmt.Errorf("db not available")
	}

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS kurls (key TEXT PRIMARY KEY, curl TEXT)")
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
