package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string
	Name   string
	Author string
}

var books = []Book{
	{ID: "1", Name: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams"},
	{ID: "2", Name: "Cloud Native Go", Author: "M.-Leander Reimer"},
	{ID: "3", Name: "Jane Eyre", Author: "Charlotte Brontë"},
}

func ListBooks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, books)
}

func GetBook(ctx *gin.Context) {
	id := ctx.Param("id")
	for _, book := range books {
		if book.ID == id {
			ctx.IndentedJSON(http.StatusOK, book)
		}
	}
}

func CreateBook(ctx *gin.Context) {
	var book Book
	if err := ctx.BindJSON(&book); err != nil {
		return
	}
	books = append(books, book)
	ctx.IndentedJSON(http.StatusCreated, book)
}

func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	for index, book := range books {
		if book.ID == id {
			var updatedBook Book
			if err := ctx.BindJSON(&updatedBook); err != nil {
				return
			}
			books = append(books[:index], books[index+1:]...)
			books = append(books, updatedBook)
			ctx.IndentedJSON(http.StatusOK, updatedBook)
		}
	}
}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			ctx.IndentedJSON(http.StatusOK, book)
		}
	}
}
