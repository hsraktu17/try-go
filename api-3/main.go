package main

import (
	"net/http"
	"strconv"
	"sync"

	"slices"

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
		v1.POST("/users", createUser)
		v1.GET("/users", getUsers)
		v1.GET("/users/:id", getUserById)
		v1.PUT("/users/:id", updatedUser)
		v1.DELETE("/users/:id", deleteUser)
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

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func getUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func updatedUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updateUser User

	if err := c.BindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, user := range users {
		if user.ID == id {
			updateUser.ID = id
			users[i] = updateUser
			c.JSON(http.StatusOK, updateUser)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

func deleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	mu.Lock()
	defer mu.Unlock()
	for i, user := range users {
		if user.ID == id {
			users = slices.Delete(users, i, i+1)
			c.JSON(http.StatusBadRequest, gin.H{"message": "user Deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}
