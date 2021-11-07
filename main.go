package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", postbooks)

	router.Run("localhost:8080")
}

type book struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Author []string `json:"author"`
	Price  float64  `json:"price"`
}

var books = []book{
	{ID: "1", Title: "Designing Web APIs: Building APIs That Developers Love", Author: []string{"Saurabh Sahni", "Brenda Jin", "Amir Shevat"}, Price: 8.99},
	{ID: "2", Title: "RESTful Web APIs: Services for a Changing World", Author: []string{"Sam Ruby", "Leonard Richardson"}, Price: 8.99},
	{ID: "3", Title: "API Design Patterns", Author: []string{"JJ Geewax"}, Price: 8.99},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func postbooks(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}
