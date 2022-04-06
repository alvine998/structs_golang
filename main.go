package main

import (
	"fmt"
	"log"
	"structs/book"
	"structs/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Start DB Connection
	dsn := "root:@tcp(127.0.0.1:3306)/structs?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}
	// Auto migrate db
	db.AutoMigrate(&book.Book{})
	// Check Print Connection
	fmt.Println("Database connection succeed")

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)

	bookHandler := handler.NewBookHandler(bookService)

	// Calling name repository
	// bookRepository := book.NewRepository(db)

	// Calling from respository to Create
	// book := book.Book{
	// 	Title:       "$100 Startup",
	// 	Description: "Good book",
	// 	Price:       95000,
	// 	Rating:      4,
	// 	Discount:    0,
	// }
	// bookRepository.Create(book)

	// Calling from repository find by id
	// book, err := bookRepository.FindByID(2)
	// fmt.Println("Title : ", book.Title)

	// Repository get all data
	// books, err := bookRepository.FindAll()
	// for _, book := range books {
	// 	fmt.Println("Title : ", book.Title)
	// }

	// End DB Connection

	// CRUD
	// =================
	// Create Data
	// =================
	// book := book.Book{}
	// book.Title = "Atomic Habits"
	// book.Price = 120000
	// book.Discount = 15
	// book.Rating = 4
	// book.Description = "Ini adalah buku yang bagus self development"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("Error creating book record")
	// }

	// =================
	// Get Data from db last
	// =================
	// var book book.Book

	// err = db.Debug().Last(&book).Error
	// if err != nil {
	// 	fmt.Println("Error finding book record")
	// }
	// fmt.Println("Title : ", book.Title)
	// fmt.Println("Book object %v", book)

	// =================
	// Get All Data from DB
	// =================
	// var books []book.Book

	// err = db.Debug().Find(&books).Error
	// if err != nil {
	// 	fmt.Println("Error finding book record")
	// }

	// for _, b := range books {
	// 	fmt.Println("Title : ", b.Title)
	// 	fmt.Println("Book object %v", b)
	// }

	// =================
	// Get Data by details using where
	// =================
	// var books []book.Book

	// err = db.Debug().Where("rating = ?", 5).Find(&books).Error
	// if err != nil {
	// 	fmt.Println("Error finding book record")
	// }

	// for _, b := range books {
	// 	fmt.Println("Title : ", b.Title)
	// 	fmt.Println("Book object %v", b)
	// }

	// =================
	// Update Data by id
	// =================
	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("Error finding book record")
	// }

	// book.Title = "Man Tiger (Revised Edition 1)"
	// err = db.Debug().Save(&book).Error
	// if err != nil {
	// 	fmt.Println("Error updating book record")
	// }

	// =================
	// Delete Data by id
	// =================
	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("Error finding book record")
	// }

	// err = db.Debug().Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("Error deleting book record")
	// }

	router := gin.Default()

	// Versioning Route
	v1 := router.Group("/v1")

	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/hello", bookHandler.HelloHandler)
	// Get by id parameter
	v1.GET("/books/:id", bookHandler.BooksHandler)
	v1.GET("/books/:id/:title", bookHandler.BooksHandler2)
	// Get by parameter query
	v1.GET("/query", bookHandler.QueryHandler)
	v1.GET("/query2", bookHandler.QueryHandler2)
	// Post Data
	v1.POST("/books", bookHandler.PostBooksHandler)

	router.Run()
}

// main layer
// handler layer
// service layer (Business Logic)
// repository layer (Berhubungan dengan db)
// db layer
// mysql layer
