package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func GetSqliteDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./monitier.db")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(50)
	return db
}
