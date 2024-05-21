package config

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var serverOrigin string = os.Getenv("SERVER_ORIGIN")

var CorsConfig gin.HandlerFunc = cors.New(cors.Config{
	AllowOrigins:     []string{serverOrigin},
	AllowMethods:     []string{"POST"},
	AllowHeaders:     []string{"Origin"},
	AllowCredentials: true,
})
