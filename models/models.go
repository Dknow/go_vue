package models

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title string `json:"title"`
	Code  string `json:"code"`
	Price uint   `json:"price"`
}

type Category struct {
	gorm.Model
	Name string
	Desc string
}

type Website struct {
	gorm.Model
	Title      string `json:"title"`
	Url        string `json:"url"`
	Categoryid uint `json:"categoryid"`
}

func initDb() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func GetProduct() (res string) {
	//	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db := initDb()
	db.AutoMigrate(&Product{}, &Category{})
	//db.Create(&Product{Title: "新款手机", Code: "D42", Price: 1000})
	//db.Create(&Product{Title: "新款电脑", Code: "D43", Price: 3500})
	//var product Product
	var products []Product
	db.Find(&products)
	r, _ := json.Marshal(products)
	//if err{
	//
	//}
	//fmt.Println(string(r))

	return string(r)
}

func GetCategory() interface{} {
	db := initDb()
	var CategoryList []Category
	db.Find(&CategoryList)
	fmt.Println(CategoryList)

	r, _ := json.Marshal(CategoryList)
	fmt.Println(r)
	//res :=string(r)
	res := make(map[string]string)
	for i:=0;i<len(CategoryList);i++{
		r[i] = json.Unmarshal([]byte(),&res)
	}
	//json.Unmarshal([]byte(r),&res)
	return res
}

func GetWebsiteList() interface{} {
	db := initDb()
	var CategoryList []Category
	db.Find(&CategoryList)
	var c_id []int
	for _,value := range CategoryList{
		c_id = append(c_id, value.ID)
	}
	fmt.Println(c_id)
	//return c_id
}