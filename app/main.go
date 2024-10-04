package main

import (
	"fmt"
	data "gosql/database"
	util "gosql/utility"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("asd", util.RandomID(10000))
	route := gin.Default()
	route.POST("/massinsert", data.MassInsert)
	route.POST("/create", data.Insert)
	route.GET("/users", data.Get)
	route.Run()

}
