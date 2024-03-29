package main

import (
	"go-pdf2jpeg/config"
	"go-pdf2jpeg/handlers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const serverPort string = ":5001"

func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	server := gin.Default()

	// Middleware
	server.Use(config.CorsConfig)

	// Routes
	server.POST("/convert", handlers.ConvertPdf)

	// Run server
	server.Run(serverPort)
}
