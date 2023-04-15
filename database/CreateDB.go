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
		email TEXT DEFAULT '',
		userKey TEXT DEFAULT ''
		);`)

	if err != nil {
		panic(err)
	}

	_, err = OpenDB().Exec(`CREATE TABLE IF NOT EXISTS wanted (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		codeCrypted TEXT DEFAULT '',
		userKey TEXT DEFAULT '',
		price INTEGER DEFAULT 0,
		);`)
	if err != nil {
		panic(err)
	}

	_, err = OpenDB().Exec(`CREATE TABLE IF NOT EXISTS sold (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		codeCrypted TEXT DEFAULT '',
		userKey TEXT DEFAULT '',
		price INTEGER DEFAULT 0,
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
