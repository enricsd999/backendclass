package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// @Summary Get user profile
// @Description Get user profile
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Security BearerAuth
// @Router /profile [get]
func HandlerUser(w http.ResponseWriter, r *http.Request) {

	//Ambil JWT Token dari header authorization
	tokenString := r.Header.Get("Authorization")

	//Hilangkan text "Bearer " dari Token String
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	} else {
		http.Error(w, "Invalid token format", http.StatusUnauthorized)
		return
	}
	//Validasi Token dan Petakan ke dalam claim
	claim, err := ValidateToken(tokenString)

	//Koneksi Database
	dsn := "root:@tcp(127.0.0.1:3306)/backend?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Buat variabel untuk menampung hasil query
	var fname, lname string

	//Gunakan claim untuk menarik data profil
	db.QueryRow(`SELECT FIRST_NAME,LAST_NAME FROM users WHERE USERNAME = ?`, claim.UserID).Scan(&fname, &lname)
	profile := struct {
		Fname string `json:"fname"`
		Lname string `json:"lname"`
	}{
		Fname: fname,
		Lname: lname,
	}
	//Tampilkan profile yang sudah terisi hasil query
	jsonData, err := json.Marshal(profile)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
