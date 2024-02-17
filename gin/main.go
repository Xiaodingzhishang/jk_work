package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.GET("/user/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "hello,这是参数路由"+name)
	})
	server.GET("/views/*.html", func(ctx *gin.Context) {
		page := ctx.Param(".html")
		ctx.String(http.StatusOK, "hello,这是通配符路由"+page)
	})
	server.GET("/order", func(ctx *gin.Context) {
		oid := ctx.Query("id")
		ctx.String(http.StatusOK, "hello,这是查询参数"+oid)
	})
	server.POST("/post", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello,post 方法")
	})
	server.Run(":8080") // listen and serve on 0.0.0.0:8080
}
