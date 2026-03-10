package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

<<<<<<< HEAD
func ConnectDatabase() *sql.DB {
=======
func ConnectDatabase() (*sql.DB, error) {
>>>>>>> empty
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal(err)
	}
<<<<<<< HEAD
	return db
=======
	return db, err
>>>>>>> empty
}
