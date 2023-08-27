package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func CreateClient() *sql.DB {
	if database != nil {
		return database
	}

	db, err := sql.Open("sqlite3", "./dev.sqlite3")

	if err != nil {
		panic(err)
	}

	database = db

	return db
}
