// Package routes
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/warmdev17/innogenlab.com/internal/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api/auth/registered", handlers.Register)

	return r
}
