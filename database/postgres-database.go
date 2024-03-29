package database

import (
	"database/sql"
	"fmt"
	"github.com/fyllekanin/com.monitier.server/config"
	_ "github.com/lib/pq"
	"log"
)

func GetPostgresDatabase(config *config.AppConfig) *sql.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		config.DatabaseUsername, config.DatabasePassword, config.DatabaseHost, config.DatabaseName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(50)
	return db
}
