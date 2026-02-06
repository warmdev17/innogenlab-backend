// Package handlers
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/warmdev17/innogenlab.com/internal/database"
	"github.com/warmdev17/innogenlab.com/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Password string
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hash),
		Role:     "student",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"err": "Email already exists"})
		return
	}

	c.JSON(200, gin.H{"msg": "User registered"})
}
