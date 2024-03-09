package utils

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var CorsConfig gin.HandlerFunc = cors.New(cors.Config{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{"POST"},
	AllowHeaders: []string{"Origin"},
})
