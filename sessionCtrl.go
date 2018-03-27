// session management

package main

import (
	"net/http"
  "os"
  "fmt"

	"github.com/gin-gonic/gin"
	)

  // Authenticate token
  func authenticateToken(c *gin.Context) {
  	token := c.GetHeader("Authorization")
  	if token == "" {
  		c.AbortWithStatusJSON(http.StatusBadRequest, &map[string](interface{}){
  			"status":  "error",
  			"code":    "2000",
  			"message": "Authorization parameters are invalid.",
  		})
  		return
  	}
    //access auth token from environment variable
    auth := os.Getenv("AUTH_TOKEN")

  	if auth != token {
  		// User does not exist
  		c.AbortWithStatusJSON(http.StatusUnauthorized, &map[string](interface{}){
  			"status":  "error",
  			"code":    "2001",
  			"message": "Invalid access-token",
  		})
  		return
  	}
  	return
  }
