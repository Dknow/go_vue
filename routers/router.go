package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_vue/routers/controllers"
)


const (
	dbDriverName = "sqlite3"
	dbName       = "./data.db3"
)


var count int32

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/count",countMiddleware, Count)
		apiRouters.GET("/GetWebsiteList", countMiddleware,controllers.GetWebsiteList)
		apiRouters.GET("/GetCategoryList",countMiddleware, controllers.GetCategoryList)

	}
}
func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/adminApi")
	{
		adminRouters.GET("/AddWebsite",countMiddleware, controllers.AddWebsite)
		adminRouters.GET("/DeleteWebsite",countMiddleware, controllers.DeleteWebsite)
		adminRouters.GET("/EditWebsite",countMiddleware, controllers.EditWebsite)
		//adminRouters.GET("/AddWebsite",countMiddleware, controllers.AddWebsite)
		//adminRouters.GET("/AddWebsite",countMiddleware, controllers.AddWebsite)

	}
}
// get website list

// admin

// 分类 crud

// website crud

// count
func countMiddleware(c *gin.Context){
	count++
	fmt.Println("count: ",count)
}

func Count(c *gin.Context) {
	c.JSON(200, gin.H{
		"count":count ,
	})
}
//	middleware

