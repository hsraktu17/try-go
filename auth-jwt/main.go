package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("supersecret")
var users = map[string]string{}

func main() {
	r := gin.Default()

	r.POST("/register", register)
	r.POST("/login")
	r.GET("/users", getAllUsers)

	auth := r.Group("/auth")
	auth.Use()
	auth.GET("/profile")

	r.Run(":8000")
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func register(c *gin.Context) {
	var cred Credentials
	if err := c.BindJSON(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid inputs"})
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(cred.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error hashing the password"})
	}

	users[cred.Email] = string(hashed)
	c.JSON(http.StatusOK, users[cred.Email])
}

func getAllUsers(c *gin.Context) {
	userList := []string{}
	for email := range users {
		userList = append(userList, email)
	}
	c.JSON(http.StatusOK, gin.H{"users": userList})
}
