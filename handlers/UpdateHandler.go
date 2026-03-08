package handlers

import (
	"encoding/json"
	"net/http"
)

func UpdateData(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Username  string `json:"username"`
		FirstName string `json:"fname"`
		LastName  string `json:"lname"`
	}{}
	json.NewDecoder(r.Body).Decode(&body)
}
