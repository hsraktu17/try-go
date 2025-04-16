package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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
		v1.GET("/users/:ID", getUserById)
		v1.PUT("/users/:ID", updateUserById)
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
	log.Println("data log", string(data))

	c.String(http.StatusOK, string(data))
}

func getUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ID"))
	log.Println("Get id params", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	data, err := os.ReadFile("user.txt")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println("Get all the user data", string(data))
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.Split(line, ", ")
		if len(parts) != 3 {
			log.Println("Malformed line:", line)
			continue
		}

		var u User
		_, err := fmt.Sscanf(parts[0], "ID: %d", &u.ID)
		if err != nil {
			log.Println("Error parsing ID:", err)
			continue
		}
		u.Name = strings.TrimPrefix(parts[1], "Name: ")
		ageStr := strings.TrimPrefix(parts[2], "Age: ")
		u.Age, err = strconv.Atoi(ageStr)

		if err != nil {
			log.Println("Error parsing user data:", line, err)
			continue
		}

		if u.ID == id {
			c.JSON(http.StatusOK, u)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

func updateUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "invalid ID"})
		return
	}

	var updatedUser User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	data, err := os.ReadFile("user.txt")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error reading the file"})
		return
	}

	lines := strings.Split(string(data), "\n")
	for _, line = range lines {

	}

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
