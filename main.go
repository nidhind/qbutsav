package main

import (
	"io"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/nidhind/qbutsav/db"
	"github.com/nidhind/qbutsav/utils"
)

var startUpTime time.Time = time.Now()
var GO_ENV=os.Getenv("GO_ENV")

func main() {

	// Initialize database
	db.InitMongo()

	// Set production variables
	if GO_ENV=="production"{
		gin.SetMode(gin.ReleaseMode)
	}

	api := gin.New()

	// Logging to a file.
	ginLogFilePath := utils.GetGinLogFilePath()
	f, _ := os.OpenFile(ginLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// Prevent redirects on trailing slashes
	api.RedirectTrailingSlash = false

	// Enable Logger
	api.Use(gin.Logger())

	// Enable CROS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AddAllowMethods("PATCH")
	api.Use(cors.New(corsConfig))

	// Mount API routes
	mountRoutes(api)

	// Default port is 8080
	// To override set PORT env variable
	api.Run()
}
