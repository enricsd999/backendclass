package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
)

type Account struct {
	Username string `json:"user"`
	Password string `json:"pass"`
}

// Login
// @Summary Log into system
// @Description Log into system
// @Tags login
// @Accept json
// @Produce json
// @Param request body Account true "User Data"
// @Success 200 {object} map[string]string
// @Router /login [post]
func Login_Attempt(w http.ResponseWriter, r *http.Request) {

	//CORS HARUS BERJALAN SEBELUM HANDLER BERJALAN
	allowedOrigins := map[string]bool{
		"http://localhost:5173": true,
	}
	origin := r.Header.Get("Origin")
	if allowedOrigins[origin] {
		w.Header().Set("Access-Control-Allow-Origin", origin) // Gunakan domain spesifik di produksi
	}
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent) // or http.StatusOK
		// w.WriteHeader(http.StatusOK) // or http.StatusOK
		return // Don't call next handler
	}

	dsn := "root:@tcp(127.0.0.1:3306)/backend?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	acc := struct {
		UsernameFromUser string `json:"user"`
		PasswordFromUser string `json:"pass"`
	}{}

	//Konversi request body (JSON) ke dalam instance
	fmt.Println("Username from User", acc.UsernameFromUser)
	err = json.NewDecoder(r.Body).Decode(&acc)
	fmt.Println("Username from User", acc.UsernameFromUser)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Lakukan query dan simpan hasil query di variabel username dan password
	var usernameFromDatabase, passwordFromDatabase, role string
	row := db.QueryRow(`SELECT USER,PASS,ROLE FROM login_account WHERE USER = ?`, acc.UsernameFromUser)
	err = row.Scan(&usernameFromDatabase, &passwordFromDatabase, &role)
	if err != nil {
		res := struct {
			Response string `json:"response"`
			Status   int    `json:"status"`
			Message  string `json:"message"`
		}{
			Response: "Unauthorized",
			Status:   401,
			Message:  "Wrong Username or Password",
		}
		jsonData, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
		fmt.Println("4", err)
		return
	}

	//Cek jika password pada request body sama dengan pada hasil query
	if acc.PasswordFromUser == passwordFromDatabase {
		token, err := GenerateJWT(usernameFromDatabase, role)
		if err != nil {
			fmt.Println(err)
		}
		res := struct {
			Response string `json:"response"`
			Code     int    `json:"code"`
			Message  string `json:"message"`
			Token    string `json:"token"`
		}{
			Response: "OK",
			Code:     200,
			Message:  "Login Success",
			Token:    token,
		}
		jsonData, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	} else {
		res := struct {
			Response string `json:"response"`
			Code     int    `json:"code"`
			Message  string `json:"message"`
		}{
			Response: "Unauthorized",
			Code:     401,
			Message:  "Wrong Username or Password",
		}
		jsonData, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}

}

func GenerateJWT(userID string, role string) (string, error) {

	var jwtKey = []byte("ini-juga-coba")
	expirationTime := time.Now().Add(15 * time.Minute) // Token expires in 15 minutes

	//Specify the claim
	claims := struct {
		UserID string `json:"user_id"`
		Role   string `json:"role"`
		jwt.RegisteredClaims
	}{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "UNKLAB-API",
			Subject:   userID,
		},
	}
	// Use HS256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
