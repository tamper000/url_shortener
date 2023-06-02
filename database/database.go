package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init() {
	db, _ = sql.Open("sqlite3", "./urls.db")

	db.Exec(CreateIfNotExists)
}

func Add(short_id, url string) {
	db.Exec(InsertNewUrl, short_id, url)
}

func Get(short_id string) string {
	var result string

	row := db.QueryRow(Select, short_id)
	row.Scan(&result)

	return result
}
