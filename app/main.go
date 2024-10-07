package main

import (
	data "gosql/database"

	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()

	//Basic Auth
	auth := route.Group("/api", gin.BasicAuth(gin.Accounts{
		"gigz": "pwpwpwpwpw",
	})).Use()

	auth.POST("/massinsert", data.MassInsert)
	auth.POST("/create", data.Insert)
	auth.GET("/user/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		data.GetDetail(id, ctx)
	})
	auth.GET("/users", data.Get)
	route.Run()

}
