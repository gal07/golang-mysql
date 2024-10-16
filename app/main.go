package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()

	// Route
	service(route)

	// Run
	route.Run()

}
