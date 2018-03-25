// Handels most of the user specific operations

package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/nidhind/qbutsav/db"
	"github.com/nidhind/qbutsav/models"
	"github.com/nidhind/qbutsav/utils"

	"time"
	"golang.org/x/crypto/bcrypt"
)

func addUserHandler(c *gin.Context) {

	// Parse request body into JSON
	var user models.SignUpReq
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &map[string](interface{}){
			"status":  "error",
			"code":    "1002",
			"message": "Error in parsing JSON input",
		})
		return
	}
	emailId := user.EmailId
	if !utils.IsValidEmail(emailId) {
		c.AbortWithStatusJSON(http.StatusBadRequest, &map[string](interface{}){
			"status":  "error",
			"code":    "1000",
			"message": "Invalid emailId",
		})
		return
	}

	if !utils.IsValidPassword(user.Password) {
		c.AbortWithStatusJSON(http.StatusBadRequest, &map[string](interface{}){
			"status":  "error",
			"code":    "1001",
			"message": "Invalid or weak password",
		})
		return
	}

	// Check if user already exists
	if DoesEmailExists(emailId) {
		c.AbortWithStatusJSON(http.StatusBadRequest, &map[string](interface{}){
			"status":  "error",
			"code":    "1003",
			"message": "User already exists",
		})
		return
	}

	// Hash the password
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hash)
	// Create User Object to insert
	u := db.InsertUserQuery{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Password:    user.Password,
		Email:       emailId,
		AccessLevel: "normal",
		Level:       1,
		PreviousLevelFinishTime: time.Now(),
	}
	err = db.InsertNewUser(&u)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &map[string](interface{}){
			"status":  "error",
			"code":    "500",
			"message": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusCreated, &map[string](interface{}){
		"status":  "success",
		"code":    "0",
		"message": "created",
	})
}

// Check if user exists by emailId
func DoesEmailExists(id string) bool {
	_, err := db.GetUserByEmail(id)
	if err != nil && err.Error() == "not found" {
		// User doesnot exists
		return false
	} else if err != nil {
		panic(err)
	}
	// User exists
	return true
}

// Fetch and serve user profile
func getUserProfile(c *gin.Context) {
	// This is a authenticated route
	// User will be already present in context
	i, _ := c.Get("user")
	u := i.(*db.User)

	r := models.ProfileRes{
		Code:   "0",
		Status: "success",
		Payload: &models.UserProfile{
			FirstName:               u.FirstName,
			LastName:                u.LastName,
			Email:                   u.Email,
			Level:                   u.Level,
			AccessLevel:             u.AccessLevel,
			PreviousLevelFinishTime: u.PreviousLevelFinishTime.String(),
		},
	}
	c.JSON(http.StatusOK, &r)
}
