package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func DeleteData(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		fmt.Println(err)
		return
	}

	body := struct {
		User string `json:"user"`
	}{}
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
		return
	}

	stmt, err := db.Prepare(`DELETE FROM login_account WHERE USER = ?`)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = stmt.Exec(body.User)
	if err != nil {
		fmt.Println(err)
		return
	}
	stmt, err = db.Prepare(`DELETE FROM users WHERE USERNAME = ?`)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = stmt.Exec(body.User)
	if err != nil {
		fmt.Println(err)
		return
	}
	response := struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		Code:    200,
		Message: "User Removed",
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
