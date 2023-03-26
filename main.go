package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID			string  `json:"id"`
    Title       string	`json:"title"`
    Author      string	`json:"author"`
    Publisher   string	`json:"publisher"`
    PublishYear int		`json:"publishYear"`
    ISBN        string	`json:"isbn"`
    Pages       int		`json:"pages"`
    Price       float64	`json:"price"`
}

var books = []Book{
	{
		ID: 		 "1",
		Title:       "The Catcher in the Rye",
		Author:      "J.D. Salinger",
		Publisher:   "Little, Brown and Company",
		PublishYear: 1951,
		ISBN:        "0316769177",
		Pages:       277,
		Price:       13.99,
	},
	{
		ID:			 "2",
		Title:       "To Kill a Mockingbird",
		Author:      "Harper Lee",
		Publisher:   "J. B. Lippincott & Co.",
		PublishYear: 1960,
		ISBN:        "0-06-112008-1",
		Pages:       281,
		Price:       12.99,
	},
	{
		ID:			 "3",
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

func addBook(context *gin.Context){
	var newBook Book

	if err := context.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)

	context.IndentedJSON(http.StatusCreated, newBook)
	
}	

func main(){
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", addBook)
	router.Run("localhost:9090")
}