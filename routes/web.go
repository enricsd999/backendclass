package routes

import (
	"loginsystem/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("POST /update", handlers.UpdateData)
	http.HandleFunc("POST /insert", handlers.InsertData)

	http.ListenAndServe(":8080", nil)
}
