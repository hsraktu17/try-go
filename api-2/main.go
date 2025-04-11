package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name" binding:"required,min=2"`
	Age  int    `json:"age" binding:"gte=18,lte=120"`
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Message": "Hello from Gin"})
	})

	router.POST("/", func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "User is valid", "user": user})
	})

	router.Run(":8000")
}
