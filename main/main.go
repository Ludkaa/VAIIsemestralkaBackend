package main

import (
	config "backend/configs"
	"backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

func main() {
	// Connect DB
	config.Connect()

	// Init Router
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "DELETE", "POST", "PUT"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Route Handlers / Endpoints
	routes.Routes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4747"
	}

	log.Fatal(router.Run(":" + port))
}
