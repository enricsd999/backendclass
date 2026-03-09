package routes

import (
	"loginsystem/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("POST /login", handlers.Login)
	http.HandleFunc("POST /update", handlers.Update)

	http.ListenAndServe(":8080", nil)
}
