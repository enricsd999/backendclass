package handlers

import (
	"fmt"
	"loginsystem/database"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/golang-jwt/jwt/v5"
)

func Login(w http.ResponseWriter, r *http.Request) {

	db, err := database.ConnectDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = db
}
