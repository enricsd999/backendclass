package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GetData(w http.ResponseWriter, r *http.Request) {

	//Batasi method pada route
	if r.Method != "POST" && r.Method != "OPTIONS" {
		http.Error(w, "Method Not Allowed", http.StatusUnauthorized)
		return
	}
	//CORS
	allowedOrigins := map[string]bool{
		"http://localhost:5173": true,
		"http://localhost:3000": true,
	}
	origin := r.Header.Get("Origin")
	if allowedOrigins[origin] {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")

	// Handle preflight requests (Actual request will ignore)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return // Jangan teruskan perintah
	}

	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := db.Query(`SELECT NAME,ROLE,DEPARTMENT,STATUS,JOINED,SALARY FROM data`)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := []struct {
		Name       string `json:"name"`
		Role       string `json:"role"`
		Department string `json:"department"`
		Status     string `json:"status"`
		Joined     string `json:"joined"`
		Salary     string `json:"salary"`
	}{}
	for rows.Next() {
		datum := struct {
			Name       string `json:"name"`
			Role       string `json:"role"`
			Department string `json:"department"`
			Status     string `json:"status"`
			Joined     string `json:"joined"`
			Salary     string `json:"salary"`
		}{}
		rows.Scan(&datum.Name, &datum.Role, &datum.Department, &datum.Status, &datum.Joined, &datum.Salary)
		data = append(data, datum)
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
