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

// Lock user by id for auction
func lockUserById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"code":    "400",
			"message": "Invalid userid",
		})
		return
	}
	// Check if user exists
	u, err := db.GetUserById(id)
	if err != nil && err.Error() == "not found" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"code":    "400",
			"message": "Invalid userid",
		})
		return
	} else if err!=nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"code":    "500",
			"message": "Internal server error",
		})
		return
	}

	//Check if user is already locked
	if u.Status == "in_progress" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"code":    "400",
			"message": "User already locked",
		})
		return
	}

	// Check if other users are already locked
	ul, err := db.GetUsersByStatus("in_progress")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"code":    "500",
			"message": "Internal server error",
		})
		return
	}
	if len(ul) != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"code":    "400",
			"message": "Some other User is already locked",
		})
		return
	}

	// Update and lock the user
	err = db.UpdateUserStatusById(id, "in_progress", 0)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"code":    "500",
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": "0",
		"message": "locked",
		"status": "success",
	})
}

// Unlock user by id for auction
func unlockUserById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"code":    "400",
			"message": "Invalid userid",
		})
		return
	}
	// Check if user exists
	u, err := db.GetUserById(id)
	if err != nil && err.Error() == "not found" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"code":    "400",
			"message": "Invalid userid",
		})
		return
	} else if err!=nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"code":    "500",
			"message": "Internal server error",
		})
		return
	}

	//Check if user is already unlocked
	if u.Status != "in_progress" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"code":    "400",
			"message": "User not locked",
		})
		return
	}

	// Update and unlock the user
	err = db.UpdateUserStatusById(id, "waiting", 0)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"code":    "500",
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": "0",
		"message": "unlocked",
		"status": "success",
	})
}