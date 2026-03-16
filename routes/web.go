package routes

import (
	"loginsystem/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("POST /insert", handlers.InsertData)
	http.HandleFunc("POST /delete", handlers.DeleteData)
	http.ListenAndServe(":8080", nil)
}
