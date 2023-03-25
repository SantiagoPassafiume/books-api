package main

type Book struct {
    Title       string `json:"title"`
    Author      string	`json:"author"`
    Publisher   string	`json:"publisher"`
    PublishYear int	`json:"publishYear"`
    ISBN        string	`json:"isbn"`
    Pages       int	`json:"pages"`
    Price       float64	`json:"price"`
}

var book = Book{
    Title:       "The Catcher in the Rye",
    Author:      "J.D. Salinger",
    Publisher:   "Little, Brown and Company",
    PublishYear: 1951,
    ISBN:        "0316769177",
    Pages:       277,
    Price:       13.99,
}