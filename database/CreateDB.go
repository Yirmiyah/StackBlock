package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDB() {

	_, err := OpenDB().Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT DEFAULT 'anonymous',
		codeCrypted TEXT DEFAULT '',
		key TEXT DEFAULT ''
		);`)

	if err != nil {
		panic(err)
	}
}

func OpenDB() *sql.DB {
	DB, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	return DB
}
