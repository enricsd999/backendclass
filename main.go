package main

import (
	"fmt"
	"loginsystem/routes"

	"github.com/joho/godotenv"
)

// @title Login API
// @version 1.0
// @description This is a description
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	routes.RegisterRoutes()
}
