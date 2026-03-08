package routes

import (
	"loginsystem/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("POST /update", handlers.UpdateData)

	http.ListenAndServe(":8080", nil)
}
