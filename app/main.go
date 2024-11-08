package main

import (
	"gosql/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load ENV
	errs := godotenv.Load(".env")
	if errs != nil {
		log.Fatal("Error loading .env file")
	}

	// server stuff
	serverPort := os.Getenv("SERVER_PORT")

	// release Mode
	// release := os.Getenv("RELEASE_MODE")
	// if release == "release" {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	route := gin.Default()
	route.Use(middleware.MiddlewareAuth)

	// Route
	service(route)

	// Run
	route.Run(serverPort)

}
