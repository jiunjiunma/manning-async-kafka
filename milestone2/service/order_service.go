package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func checkHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func main() {
	router := gin.Default()
	router.GET("/health", checkHealth)

	router.Run("localhost:8080")
}
