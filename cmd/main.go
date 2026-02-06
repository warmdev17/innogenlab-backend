// Package main
package main

import (
	"log"

	"github.com/warmdev17/innogenlab.com/internal/database"
	"github.com/warmdev17/innogenlab.com/internal/routes"
)

func main() {
	database.Connect()
	r := routes.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server start failed:", err)
	}
}
