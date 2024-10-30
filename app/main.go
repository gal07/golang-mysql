package main

import (
	"gosql/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// server stuff
	serverPort := os.Getenv("SERVER_PORT")

	route := gin.Default()
	route.Use(middleware.MiddlewareAuth)

	// Route
	service(route)

	// Run
	route.Run(serverPort)

}
