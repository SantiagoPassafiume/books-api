package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
    Title       string	`json:"title"`
    Author      string	`json:"author"`
    Publisher   string	`json:"publisher"`
    PublishYear int	`json:"publishYear"`
    ISBN        string	`json:"isbn"`
    Pages       int	`json:"pages"`
    Price       float64	`json:"price"`
}

var books = []Book{
	{
		Title:       "The Catcher in the Rye",
		Author:      "J.D. Salinger",
		Publisher:   "Little, Brown and Company",
		PublishYear: 1951,
		ISBN:        "0316769177",
		Pages:       277,
		Price:       13.99,
	},
	{
		Title:       "To Kill a Mockingbird",
		Author:      "Harper Lee",
		Publisher:   "J. B. Lippincott & Co.",
		PublishYear: 1960,
		ISBN:        "0-06-112008-1",
		Pages:       281,
		Price:       12.99,
	},
	{
		Title:       "1984",
		Author:      "George Orwell",
		Publisher:   "Secker & Warburg",
		PublishYear: 1949,
		ISBN:        "978-0-14-118776-1",
		Pages:       328,
		Price:       8.99,
	},
}

func getBooks(context *gin.Context){
	context.IndentedJSON(http.StatusOK, books)

}
