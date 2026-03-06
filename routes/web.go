package routes

import (
	"loginsystem/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/login", handlers.Login_Attempt)
	http.HandleFunc(" /profile", handlers.HandlerUser)
	http.HandleFunc("/", handlers.HandlerData)

	http.ListenAndServe(":8080", nil)
}
