package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jiunjiunma/manning-async-kafka/milestone3/model"
	"github.com/jiunjiunma/manning-async-kafka/milestone3/sender"
	"net/http"
)



func postOrder(c *gin.Context, s sender.OrderSender) {
	var order model.Order

	// Call BindJSON to bind the received JSON to Order .
	if err := c.BindJSON(&order); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	err := s.SendOrder(order)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to send order"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"Message": "Order successfully sent"})
}


func main() {
	s, err := sender.NewSender()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	router := gin.Default()

	router.POST("/orders", func(c *gin.Context) {
		postOrder(c, s)
	})
	router.GET("/health", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "OK"})
	})

	router.Run("localhost:8080")
}
