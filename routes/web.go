package routes

import (
	"loginsystem/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/getdata", handlers.GetData)
	http.ListenAndServe(":8080", nil)
}
