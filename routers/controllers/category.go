package controllers

import (
	"github.com/gin-gonic/gin"
	"go_vue/models"
)

func GetCategoryList(c *gin.Context) {
	res := models.GetCategory()
	c.JSON(200, gin.H{
		"message": "GetCategoryList",
		"data":res,
	})
}