package routes

import (
	"databaseconnect/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/", handlers.HandlerData)
	// http.Handle("/", middlewares.CheckAPI(handlers.HandlerData))
	http.ListenAndServe(":8080", nil)
}
