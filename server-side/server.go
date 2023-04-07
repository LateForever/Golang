package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// @Function: Create a new gin router
	router = gin.Default()

	// @Function: Load all the templates from the templates folder
	router.LoadHTMLGlob("templates/*")

	// @Router / [GET] - Homepage route
	router.GET("/", func(c * gin.Context) {
		// @Function: Render the homepage.html template with the given title and description variables
		c.HTML(http.StatusOK, "homepage.html", gin.H{
			"title": "Main website",
			"title_home": "Home",
			"description": "This is the main website",
		})
	})

	// @Router /producst [GET] - products route
	router.GET("/products", func(c * gin.Context) {
		// @Function: Render the products.html template with the given title and description variables
		c.HTML(http.StatusOK, "products.html", gin.H{
			"title": "Products",
			"title_products": "Products",
			"products": []string{"Product 1", "Product 2", "Product 3"},
		})
	})

	// @Function: Run the server on port 8080
	router.Run("192.168.31.150:8080")
}