package main

import (
	"fmt"
	"loginsystem/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	routes.RegisterRoutes()
}
