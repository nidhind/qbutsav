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
	app.POST("/teams",authenticateToken, createNewTeam)
	// Allocate user to a team
	app.PUT("/teams/user", authenticateToken,allocateUserToTeam)

	// Get users list
	app.GET("/users",getUserProfiles)
	// Lock user for auction
	app.GET("/users/lock/:id",authenticateToken, lockUserById)
	// Un-lock user for auction
	app.DELETE("/users/lock/:id",authenticateToken, unlockUserById)

	// Handle 404
	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, map[string](string){
			"message": "Resource not found",
		})
	})
}
