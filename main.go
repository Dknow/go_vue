package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_vue/routers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title  string
	Code  string
	Price uint
}
//type Category struct {
//	gorm.Model
//	Id	int16
//	Name string

//}


func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	var count int
	fmt.Println("count： ",count)

	//db

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})
	//db.Create(&Product{Title: "新款手机", Code: "D42", Price: 1000})
	//db.Create(&Product{Title: "新款电脑", Code: "D43", Price: 3500})
	//var product Product
	var products []Product

	//db.First(&product) // find product with integer primary key
	//fmt.Println(product.Title)

	db.Find(&products)
	fmt.Println(products.Ti)

	//db.First(&product, "code = ?", "D42") // find product with code D42
	//fmt.Println(product.Code)

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