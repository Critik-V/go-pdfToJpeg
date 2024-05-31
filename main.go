package main

import (
	"go-pdf2jpeg/handlers"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const serverPort string = ":5001"

func init() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

func main() {
	if os.Getenv("GIN_MODE") != "" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()

	// Middleware
	var serverOrigin string = os.Getenv("SERVER_SERVICE_ORIGIN")
	if serverOrigin == "" {
		serverOrigin = "*"
	}
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{serverOrigin},
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Origin"},
		AllowCredentials: true,
	}))

	// Routes
	server.POST("/convert", handlers.ConvertPdf)

	// Run server
	server.Run(serverPort)
}
