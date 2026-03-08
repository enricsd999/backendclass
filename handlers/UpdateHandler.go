package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func UpdateData(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Username  string `json:"username"`
		FirstName string `json:"fname"`
		LastName  string `json:"lname"`
	}{}
	json.NewDecoder(r.Body).Decode(&body)
	_, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		fmt.Println(err)
		return
	}
}
