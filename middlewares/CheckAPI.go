package middlewares

import (
	"fmt"
	"net/http"
	"os"
)

func CheckAPI(route http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		validKey := os.Getenv("API_KEY")
		apiKey := r.Header.Get("X-API-Key")

		if apiKey != validKey {
			fmt.Println("Someone trying to access your API")
			fmt.Fprintf(w, "Unauthorized Access")
			return
		}
		route.ServeHTTP(w, r)
	}
}
