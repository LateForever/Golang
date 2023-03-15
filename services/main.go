package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// We use *gin.Context instead of gin.Context because gin.Context
// is a struct and can be quite large in size. Passing it as a value
// parameter to a function will create a copy of the struct, which
// can be expensive in terms of memory and processing time.

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}