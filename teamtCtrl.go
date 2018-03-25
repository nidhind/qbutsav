package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nidhind/qbutsav/db"
	"github.com/nidhind/qbutsav/models"
)

// Fetch and serve user profile
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

	for _, v := range t {
		tr=append(tr, models.TeamProfile{
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
		Payload: tr,
	}
	c.JSON(http.StatusOK, &r)
}

