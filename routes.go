package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func mountRoutes(app *gin.Engine) {

	// Get server status
	app.GET("/status", statusHandler)

	// Get user profile
	app.GET("/teams", getTeamProfiles)
	// Add new user
	app.POST("/teams", createNewTeam)

	// Handle 404
	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, map[string](string){
			"message": "Resource not found",
		})
	})
}