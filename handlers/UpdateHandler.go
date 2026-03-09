package handlers

import (
	"loginsystem/database"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDatabase()
	_ = db
}
