package routes

import (
	"loginsystem/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/update", handlers.UpdateData)

	http.ListenAndServe(":8080", nil)
}
