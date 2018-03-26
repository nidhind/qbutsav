package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nidhind/qbutsav/db"
	"github.com/nidhind/qbutsav/models"
	"time"
	"fmt"
	"errors"
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
			Points:v.Points,
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

// Allocate user to a team
func allocateUserToTeam(c *gin.Context) {
	var allocateReq models.AllocateUserReq
	err := c.ShouldBindJSON(&allocateReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"code":    "1002",
			"message": "Error in parsing JSON input",
		})
		return
	}
	id := allocateReq.UserId
	team := allocateReq.TeamId
	points:=allocateReq.Points

	// Check if user exists
	u, err := db.GetUserById(id)
	if err != nil && err.Error() == "not found" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"code":    "400",
			"message": "Invalid userid",
		})
		return
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"code":    "500",
			"message": "Internal server error",
		})
		return
	}

	//Check if user is locked
	if u.Status != "in_progress" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"code":    "1002",
			"message": "User is not locked",
		})
		return
	}

	// Check if team exists
	t, err := db.GetTeamById(team)
	if err != nil && err.Error() == "not found" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"code":    "400",
			"message": "Invalid team id",
		})
		return
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"code":    "500",
			"message": "Internal server error",
		})
		return
	}

	// Check if user is already allocated
	err = nil
	for _, au := range t.AccquiredMembers {
		if au.Id == u.Id {
			err = errors.New("already_allocated")
		}
		break
	}
	if err!=nil && err.Error() == "already_allocated" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"code":    "400",
			"message": "User already allocated",
		})
		return
	}

	// Set new team points
	t.Points = t.Points - points
	// Add user to team
	t.AccquiredMembers = append(t.AccquiredMembers, db.TeamMembers{
		Id:u.Id,
		FirstName:u.FirstName,
		LastName:u.LastName,
		Email:u.Email,
		Image:u.Image,
		UpdatedAt:time.Now().Unix(),
	})
	// Update Teams
	err = db.UpdateTeamById(t)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"code":    "500",
			"message": "Internal server error",
		})
		return
	}

	// Update and unlock the user
	err = db.UpdateUserStatusById(u.Id, "done", points)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"code":    "500",
			"message": "Internal server error",
		})
		return
	}

	ah:=db.AuctionHistory{
		UserID:u.Id,
		TeamID:t.ID,
		TeamPoints:t.Points,
		UserPoints:points,
		At:time.Now().Unix(),
	}
	go db.InsertAuctionHistory(&ah)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"code":    "0",
		"message": "Allocated",
	})
}