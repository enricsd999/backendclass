package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	//Jangan lupa lakukan "go mod init nama_project_anda"
	//Jangan lupa lakukan "go get github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
)

type Customers struct {
	ID   int
	Name string
}

func HandlerData(w http.ResponseWriter, r *http.Request) {

	//Define Data Source Name
	//Username: root
	//Password: passcode (Jika tidak ada, bisa dikosongkan)
	//Database IP: 127.0.0.1
	//Database Port: 3306
	//Database Name: latihan

	tokenString := r.Header.Get("Authorization")
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	} else {
		http.Error(w, "Invalid token format", http.StatusUnauthorized)
		return
	}
	claim, err := ValidateToken(tokenString)
	_ = claim
	if err != nil {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}
	dsn := os.Getenv("DSN")

	//Connect to Database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Check Connection
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := db.Query(`SELECT id,name FROM cust`)
	if err != nil {
		fmt.Println(err)
		return
	}

	var customers []Customers
	for res.Next() {
		var customer Customers
		res.Scan(&customer.ID, &customer.Name)
		customers = append(customers, customer)
	}

	jsonData, err := json.Marshal(customers)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

	// fmt.Fprintf(w, "Connection Success")

	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// apiKey := r.Header.Get("X-API-Key")

	// if apiKey != os.Getenv("API_KEY") {
	// 	http.Error(w, "Unauthorized: Invalid API key", http.StatusUnauthorized)
	// 	return
	// }

}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Validate the expected signing method is HMAC (for HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// return handlers.JwtKey, nil
		var JwtKey = []byte(os.Getenv("JWT_KEY"))
		return JwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
