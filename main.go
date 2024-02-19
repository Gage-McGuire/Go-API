package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// acts as the database for the API
var books = []book{
	{ID: "1", Title: "Book One", Author: "Author One", Quantity: 2},
	{ID: "2", Title: "Book Two", Author: "Author Two", Quantity: 1},
	{ID: "3", Title: "Book Three", Author: "Author Three", Quantity: 2},
	{ID: "4", Title: "Book Four", Author: "Author Four", Quantity: 5},
	{ID: "5", Title: "Book Five", Author: "Author Five", Quantity: 3},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func addBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", addBook)
	router.GET("/books/:id", bookById)
	router.Run("localhost:8080")
}
