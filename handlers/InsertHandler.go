package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func InsertData(w http.ResponseWriter, r *http.Request) {
	body := struct {
		ProductName string `json:"product_name"`
		Owner       string `json:"owner"`
	}{}
	json.NewDecoder(r.Body).Decode(&body)
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		fmt.Println(err)
		return
	}
	stmt, err := db.Prepare(`INSERT INTO products (PRODUCT_NAME,OWNER) VALUES (?,?)`)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = stmt.Exec(body.ProductName, body.Owner)
	if err != nil {
		fmt.Println(err)
		return
	}
	res := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		Code:    "200",
		Message: "Data Inserted",
	}
	jsonData, err := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}
