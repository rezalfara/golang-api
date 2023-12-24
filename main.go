package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db connection error")
	}

	db.AutoMigrate(&book.Book{})
	//CRUD

	//===========
	//Create Data
	//===========

	// book := book.Book{}
	// book.Title = "Buku baru"
	// book.Price = 5000
	// book.Discount = 80
	// book.Rating = 3
	// book.Description = "Ini adalah buku baru"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("==========================")
	// }

	//Read data
	var book book.Book
	err = db.Debug().Where("id = ?", 1).First(&book).Error
	if err != nil {
		fmt.Println("==========================")
		fmt.Println("Error finding book record")
		fmt.Println("==========================")
	}

	//Update data
	// book.Title = "Man Tiger (Revised Edition)"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error updating book record")
	// 	fmt.Println("==========================")
	// }

	//Delete data
	err = db.Delete(&book).Error
	if err != nil {
		fmt.Println("==========================")
		fmt.Println("Error deleting book record")
		fmt.Println("==========================")
	}

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run()
}
