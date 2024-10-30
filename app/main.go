package main

import (
	"gosql/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	route := gin.Default()
	route.Use(middleware.MiddlewareAuth)

	// Route
	service(route)

	// Run
	route.Run()

}
