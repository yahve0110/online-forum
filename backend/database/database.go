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
		nickname TEXT NOT NULL UNIQUE,
		age INTEGER NOT NULL,
		gender TEXT NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = Db.Exec(`
		CREATE TABLE IF NOT EXISTS sessions (
			token TEXT PRIMARY KEY,
			user_id TEXT
		)
	`)
	if err != nil {
		log.Fatal(err)
	}


_, err = Db.Exec(`
	CREATE TABLE IF NOT EXISTS posts (
		id TEXT PRIMARY KEY,
		user_nickname TEXT NOT NULL,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		creation_time DATETIME NOT NULL,
		FOREIGN KEY (user_nickname) REFERENCES users(nickname)
	);
`)
if err != nil {
	log.Fatal(err)
}

_, err = Db.Exec(`
CREATE TABLE IF NOT EXISTS comments(
		id TEXT PRIMARY KEY,
		post_id TEXT NOT NULL,
		author TEXT NOT NULL,
		content TEXT NOT NULL,
		creation_time DATETIME NOT NULL,
		FOREIGN KEY (post_id) REFERENCES posts(id),
		FOREIGN KEY (author) REFERENCES users(nickname)
);
`)
if err != nil {
	log.Fatal(err)
}

}


