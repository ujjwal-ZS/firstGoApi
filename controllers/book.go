package controllers

import (
	"firstProject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /books
// Get all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	println(books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}
