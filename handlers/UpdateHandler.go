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
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
		return
	}

	dsn := os.Getenv("DSN")

	//Connect to Database
	db, err := sql.Open("mysql", dsn)
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
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		Code:    200,
		Message: "Data Updated",
	}

	jsonData, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
