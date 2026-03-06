package main

import (
	"databaseconnect/routes"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
	routes.RegisterRoutes()
}
