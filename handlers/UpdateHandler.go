package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func UpdateData(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Username  string `json:"username"`
		FirstName string `json:"fname"`
		LastName  string `json:"lname"`
	}{}
	json.NewDecoder(r.Body).Decode(&body)
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		fmt.Println(err)
		return
	}
	stmt, err := db.Prepare(`UPDATE users SET FIRST_NAME = ? , LAST_NAME = ? WHERE USERNAME = ?`)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = stmt.Exec(body.FirstName, body.LastName, body.Username)
	if err != nil {
		fmt.Println(err)
		return
	}
	res := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}
	jsonData, err := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}
