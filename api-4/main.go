package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required,min=2"`
	Age  int    `json:"age" binding:"required,gte=18,lte=100"`
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
		v1.POST("/users", userCreate)
		v1.GET("/users", getUsers)
	}

	router.Run(":8000")
}

func userCreate(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	mu.Lock()
	user.ID = nextID
	nextID++
	users = append(users, user)
	mu.Unlock()

	entry := fmt.Sprintf("ID: %d, Name: %s, Age: %d", user.ID, user.Name, user.Age)
	err := appendToFile("user.txt", entry)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write to file"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func getUsers(c *gin.Context) {
	data, err := os.ReadFile("user.txt")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.String(http.StatusOK, string(data))
}

func appendToFile(filename, text string) error {
	if text[len(text)-1] != '\n' {
		text += "\n"
	}
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(text); err != nil {
		return err
	}
	return nil
}
