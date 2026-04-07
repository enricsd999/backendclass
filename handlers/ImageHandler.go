package handlers

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func imageURLToBase64(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded, nil
}
func GetImage(w http.ResponseWriter, r *http.Request) {

	dsn := os.Getenv("DSN")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}

	body := struct {
		Name string `json:"name"`
	}{}

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var url string
	err = db.QueryRow(`SELECT URL FROM images WHERE NAME = ?`, body.Name).Scan(&url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fullURL := "./uploads/" + url

	base64Str, err := imageURLToBase64(fullURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	w.Write([]byte(base64Str))
}
