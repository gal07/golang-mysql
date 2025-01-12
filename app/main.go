package main

import (
	"gosql/middleware"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
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
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS", "DELETE", "PATCH"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour
	route.Use(cors.New(config))
	route.Use(middleware.MiddlewareAuth)

	// Route
	service(route)

	// Run
	route.Use()
	route.Run(serverPort)

}
