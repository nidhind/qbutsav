// Shows simple server status
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func statusHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"name":   "QBUtsav",
		"env":    os.Getenv("GO_ENV"),
		"uptime": fmt.Sprint(time.Since(startUpTime)),
	})
}
