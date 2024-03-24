package main

import (
	"go-pdf2jpeg/handlers"
	"go-pdf2jpeg/utils"
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
	server.Use(utils.CorsConfig)

	// Routes
	server.POST("/convert", handlers.ConvertPdf)

	// Run server
	server.Run(serverPort)
}

// curl -X POST http://localhost:5001/convert -d '{"fileName": "test"}' -H "Content-Type: application/json"
