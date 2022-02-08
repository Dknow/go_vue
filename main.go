package main

import (
	"github.com/gin-gonic/gin"
	"go_vue/routers"
)


func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	var count int
	r.GET("/", func(c *gin.Context) {
		//输出json结果给调用方
		count ++
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	routers.ApiRoutersInit(r)
	routers.AdminRoutersInit(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}