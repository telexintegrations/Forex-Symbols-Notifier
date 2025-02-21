package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/telexintegrations/Forex-Symbols-Notifier/routes"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		// log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server:= gin.Default()
	config := cors.Config{
		AllowOrigins:     []string{fmt.Sprintf("http://localhost:%s", port), "https://telex.im"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	// Apply CORS middleware
	server.Use(cors.New(config))

	routes.RegisterRoutes(server)

	server.Run(":8080")

}