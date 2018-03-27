package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/nidhind/qbutsav/db"
	"github.com/nidhind/qbutsav/models"
)

// Fetch and serve teams profile
func tickerHandler(c *gin.Context) {
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

	r := models.TickerRes{}
	r.Payload.TeamList = tr

	// Fetch current auction user
	u, err := db.GetUsersByStatus("in_progress")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"code":    "500",
			"message": "Internal server error",
		})
		return
	}
if len(u)>0{
	r.Payload.CurrentAuction.ID = u[0].Id
	r.Payload.CurrentAuction.FirstName = u[0].FirstName
	r.Payload.CurrentAuction.LastName = u[0].LastName
	r.Payload.CurrentAuction.Email = u[0].Email
	r.Payload.CurrentAuction.Image = u[0].Image
	r.Payload.CurrentAuction.Status = u[0].Status
	r.Payload.CurrentAuction.UpdatedAt = u[0].UpdatedAt

}
	// Fetch audtion history
	ah, err := db.FetchAuctionHistory()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"code":    "500",
			"message": "Internal server error",
		})
		return
	}
	for _, v := range *ah {
		t := models.AuctionHistoryRes{}
		t.UserID = v.UserID
		t.UserName=v.UserName
		t.TeamName=v.TeamName
		t.TeamID = v.TeamID
		t.At=v.At
		r.Payload.AuctionHistory = append(r.Payload.AuctionHistory, t)
	}

	// Get total user count
	count, err:= db.GetUserCount()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"code":    "500",
			"message": "Internal server error",
		})
		return
	}
	r.Payload.TotalUsers=count

	// Get auctioned
	au:=0
	for _,v:=range *t{
		au=au+len(v.AccquiredMembers)
	}
	r.Payload.AuctionedUsers=au

	c.JSON(http.StatusOK, &r)
}
