package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"requied,min=2"`
	Age  int    `json:"age" binding:"required,get=18,lte=100"`
}

var (
	users  []User
	nextID = 1
	mu     sync.Mutex
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/users", createUser)
	}

	router.Run(":8000")
}

func createUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	mu.Lock()
	user.ID = nextID
	nextID++
	users = append(users, user)
	mu.Unlock()

	c.JSON(http.StatusCreated, user)
}
