package postgres

import (
	"database/sql"
	"log"

	_ "github.com/go-pg/pg/v10"
)

func NewPostgresDB(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}