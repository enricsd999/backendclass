package handlers

import (
	"loginsystem/database"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Login(w http.ResponseWriter, r *http.Request) {

	db := database.ConnectDatabase()
	_ = db
}
