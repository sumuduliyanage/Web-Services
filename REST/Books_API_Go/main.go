package main

import (
	"github.com/book_service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/books/list", handlers.ListBooks)
	router.GET("/books/:id", handlers.GetBook)
	router.POST("/books", handlers.CreateBook)
	router.PUT("/books/:id", handlers.UpdateBook)
	router.DELETE("/books/:id", handlers.DeleteBook)

	router.Run(":8080")
}
