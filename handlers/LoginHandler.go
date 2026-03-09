package handlers

import (
	"fmt"
	"loginsystem/database"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Login(w http.ResponseWriter, r *http.Request) {

	db, err := database.ConnectDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = db
}
