package controllers

import (
	"KK/KKniverse-GO/example/logrokcet/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var books = []models.Book{}

func GetAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBooks := models.Book{
		ID:     len(books) + 1,
		Title:  input.Title,
		Author: input.Author,
	}
	books = append(books, newBooks)

	c.JSON(http.StatusOK, newBooks)
}
