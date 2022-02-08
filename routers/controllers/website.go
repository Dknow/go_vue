package controllers

import "github.com/gin-gonic/gin"

// GetWebsiteList
func GetWebsiteList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "weblist",
	})
}

func AddWebsite(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "AddWebsite",
	})
}
func DeleteWebsite(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DeleteWebsite",
	})
}
func EditWebsite(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "EditWebsite",
	})
}