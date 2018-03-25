package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func mountRoutes(app *gin.Engine) {

	// Get server status
	app.GET("/status", statusHandler)

	// Get team profile
	app.GET("/teams", getTeamProfiles)
	// Add new team
	app.POST("/teams", createNewTeam)

	// Get users list
	app.GET("/users",getUserProfiles)

	// Handle 404
	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, map[string](string){
			"message": "Resource not found",
		})
	})
}
