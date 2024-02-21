package handler

import (
	"firstProject/controllers"
	"firstProject/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")

	gin.SetMode(gin.ReleaseMode) // Set Gin to production mode
	router := gin.New()          // Create a new Gin router
	router.Use(gin.Logger())     // Use the default Logger middleware
	router.Use(gin.Recovery())   // Use the default Recovery middleware

	models.ConnectDatabase()

	router.GET("/books", controllers.FindBooks) // new
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	// Instead of r.Run(), use router.ServeHTTP(w, r) to handle the request
	router.ServeHTTP(w, r)
}
