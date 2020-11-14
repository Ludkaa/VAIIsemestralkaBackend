package main

import (
	config "backend/configs"
	"backend/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	// Connect DB
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	log.Fatal(router.Run(":" + port))
}
