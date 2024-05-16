package database

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

var db *sql.DB

func InitDB() error {
	database, err := sql.Open("sqlite", "outfits.db")
	if err != nil {
		return err
	}
	db = database
	return nil
}

func CloseDB() {
	db.Close()
}

func CheckConnection() error {
	pingErr := db.Ping() // confirm that connecting to the database works
	if pingErr != nil {
		return pingErr
	}
	return nil
}

func GetDB() *sql.DB {
	return db
}
