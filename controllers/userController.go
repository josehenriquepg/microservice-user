package controllers

import (
	"net/http"
	"user-management-api/config"
	"user-management-api/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("2RBF/4lD1wBMWxiCWjV4gB7aLB4Rg41Rv4S1+w6iixk=")

type Credentials struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

type Claims struct {
	Username string `json: "username"`
	jwt.StandardClaims
}

// User registration
func Register(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
		return
	}

	user := models.User{Username: creds.Username, Password: string(hashedPassword)}
	db := config.GetDB()
	_, err = user.SaveUser(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving user"})
		return
	}

	c.JSON((http.StatusCreated, gin.H{"message": "User registered sucessfully"}))
}

