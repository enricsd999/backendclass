package main

import (
	"backendclass/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/data", handlers.FetchData)
	http.ListenAndServe(":8080", nil)

	fmt.Println("Server running at http://localhost:8080")
}
