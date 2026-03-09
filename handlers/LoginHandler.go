package handlers

import (
	"loginsystem/database"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDatabase()
	_ = db
}
