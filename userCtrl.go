package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/nidhind/qbutsav/db"
	"github.com/nidhind/qbutsav/models"
)


// Fetch and serve user profile
func getUserProfiles(c *gin.Context) {
	// Fetch teams from DB

	sp := c.DefaultQuery("points", "false")

	ul, err := db.GetAllUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"code":    "500",
			"message": "Internal server error",
		})
		return
	}

	ulr := []models.UserProfile{}

	for _, v := range *ul {
		t := models.UserProfile{
			Id:v.Id,
			FirstName:v.FirstName,
			LastName:v.LastName,
			Email:  v.Email,
			Image:v.Image,
			AccessLevel:v.AccessLevel,
			Status:v.Status,
			UpdatedAt:v.UpdatedAt,
		}

		//TODO: Validate permission also
		if sp == "true" {
			t.Points = v.Points
		}
		ulr = append(ulr, t)

	}

	r := models.UserListRes{
		Code:   "0",
		Status: "success",
		Payload: &ulr,
	}
	c.JSON(http.StatusOK, &r)
}
