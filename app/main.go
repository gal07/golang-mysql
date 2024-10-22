package main

import (
	"gosql/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()
	route.Use(middleware.MiddlewareAuth)

	// Route
	service(route)

	// Run
	route.Run()

}
