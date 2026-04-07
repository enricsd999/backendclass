package routes

import (
	"loginsystem/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/upload", handlers.UploadFile)
	http.HandleFunc("/getimage", handlers.GetImage)
	http.ListenAndServe(":8080", nil)
}
