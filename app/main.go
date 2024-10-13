package main

import (
	data "gosql/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()

	//Basic Auth
	auth := route.Group("/api", gin.BasicAuth(gin.Accounts{
		"myuser": "pass123",
	})).Use()

	auth.POST("/user/massinsert", data.MassInsert)
	auth.POST("/user/create", data.Insert)
	auth.POST("/user/search", data.Search)
	auth.POST("/users", data.Get)

	auth.GET("/user/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		data.GetDetail(id, ctx)
	})

	auth.PUT("/user/update/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		is, err := strconv.Atoi(id)
		if err != nil {

			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			panic(err)

		} else {
			data.Update(is, ctx)
		}
	})

	auth.DELETE("/user/delete/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		is, err := strconv.Atoi(id)
		if err != nil {

			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			panic(err)

		} else {
			data.Delete(is, ctx)
		}
	})

	route.Run()

}
