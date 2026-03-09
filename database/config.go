package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
