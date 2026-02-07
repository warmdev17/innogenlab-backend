// Package middleware
package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("golangisthebestbackendlanguageintheworld")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.JSON(401, gin.H{"error": "No token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(auth, "Bearer ")

		token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if token == nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Next()
	}
}
