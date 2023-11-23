package database

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3" 
)

var Db *sql.DB

func InitDB() {
	var err error
	Db, err = sql.Open("sqlite3", "./database/forum.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = Db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			nickname TEXT,
			age TEXT,
			gender TEXT,
			first_name TEXT,
			last_name TEXT,
			email TEXT,
			password TEXT
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	if Db != nil {
		Db.Close()
	}
}
