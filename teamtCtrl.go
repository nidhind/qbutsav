package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nidhind/qbutsav/db"
	"github.com/nidhind/qbutsav/models"
	"time"
	"fmt"
)

// Fetch and serve teams profile
func getTeamProfiles(c *gin.Context) {
	// Fetch teams from DB

	t, err := db.GetAllTeams()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"code":    "500",
			"message": "Internal server error",
		})
		return
	}

	tr := []models.TeamProfile{}

	for _, v := range *t {
		tr = append(tr, models.TeamProfile{
			ID : v.ID,
			Name : v.Name,
			Captain : v.Captain,
			Owner : v.Owner,
			AccquiredMembers : v.AccquiredMembers,
			RelievedMembers : v.RelievedMembers,
		})
	}

	r := models.TeamProfileRes{
		Code:   "0",
		Status: "success",
		Payload: &tr,
	}
	c.JSON(http.StatusOK, &r)
}

// Create new team
func createNewTeam(c *gin.Context) {

	//TODO: Validate permission before continuing

	// Parse request body into JSON
	var team models.NewTeam
	err := c.ShouldBindJSON(&team)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"code":    "1002",
			"message": "Error in parsing JSON input",
		})
		return
	}

	//TODO: Validate the team input data

	// Generate ID - use nano time for easiness :D
	id := fmt.Sprintf("%d", time.Now().UnixNano())

	// Create team Object to insert
	t := db.Team{
		Name:team.Name,
		ID:id,
		Captain:team.Captain,
		Owner:team.Owner,
		Points:team.Points,
	}
	err = db.CreateNewTeam(&t)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"code":    "500",
			"message": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"code":    "0",
		"message": "created",
	})

}
