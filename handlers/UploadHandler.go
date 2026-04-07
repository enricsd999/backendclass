package handlers

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {

	// Only allow POST
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Batasi ukuran file (max 10MB)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get file from form field "file"
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	//Checklist
	//Batasi tipe file
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		http.Error(w, "Cannot read file", http.StatusInternalServerError)
		return
	}

	filetype := http.DetectContentType(buffer)
	fmt.Println("Type:", filetype)

	allowed := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
	}

	if !allowed[filetype] {
		http.Error(w, "Invalid file type", http.StatusBadRequest)
		return
	}
	ext := ""
	if filetype == "image/png" {
		ext = ".png"
	}
	if filetype == "image/jpeg" {
		ext = ".jpeg"
	}

	// Reset file pointer
	file.Seek(0, 0)

	//Checklist
	//Batasi penamaan file
	now := time.Now().Format("2006-01-02")
	filename := now + ext
	fmt.Println("Uploaded File:", filename)
	fmt.Println("File Size:", header.Size)
	fmt.Println("MIME Header:", header.Header)

	// Create destination file
	dst, err := os.Create("./uploads/" + filename)
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy uploaded file to destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	dsn := os.Getenv("DSN")

	//Connect to Database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	stmt, err := db.Prepare(`INSERT INTO images (NAME,URL) VALUES (?,?)`)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = stmt.Exec("Gambar", filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully")
}
