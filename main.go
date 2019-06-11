package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type(
	todoModel struct{
		gorm.Model
		Title		string	`json:"title"`
		Completed 	int    	`json:"completed"`
	}
	transformedTodo struct{
		ID        	uint   	`json:"id"`
  		Title     	string 	`json:"title"`
  		Completed 	bool   	`json:"completed"`
	}
)

func init()  {
	var err error
	db, err = gorm.Open("mysql", "root:@/api-todo?charset=utf8&parseTime=True&loc=Local")

	if err != nil{
		panic("failed to connect database")
	}
	db.AutoMigrate(&todoModel{})
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1/todos"){
		v1.POST("/", createTodo)
		v1.GET("/", fetchAllTodo)
		v1.GET("/:id", fetchSingleTodo)
		v1.PUT("/:id", updateTodo)
		v1.DELETE("/:id", deleteTodo)
	}
	router.Run()
}
