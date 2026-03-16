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

	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		fmt.Println(err)
		return
	}

	body := struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		User        string `json:"user"`
		Pass        string `json:"pass"`
		ConfirmPass string `json:"confirm_pass"`
		Email       string `json:"email"`
		DOB         string `json:"dob"`
		Role        string `json:"role"`
	}{}
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
		return
	}

	if body.Pass != body.ConfirmPass {
		response := struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    401,
			Message: "Passwords do not match",
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
		return
	}

	stmt, err := db.Prepare(`INSERT INTO login_account (USER, PASS, ROLE) values (?,?,?)`)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = stmt.Exec(body.User, body.Pass, body.Role)
	if err != nil {
		fmt.Println(err)
		return
	}
	stmt, err = db.Prepare(`INSERT INTO users (USERNAME, FNAME, LNAME, EMAIL_ADDR, DOB) values (?,?,?,?,?)`)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = stmt.Exec(body.User, body.FirstName, body.LastName, body.Email, body.DOB)
	if err != nil {
		fmt.Println(err)
		return
	}
	response := struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		Code:    200,
		Message: "User Created",
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
