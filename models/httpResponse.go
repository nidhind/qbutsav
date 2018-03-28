package models

import "github.com/nidhind/qbutsav/db"

type LoginRes struct {
	Code    string `json:"code"`
	Status  string `json:"status"`
	Payload struct {
				AccessToken string `json:"access_token"`
			}
}

type TeamProfileRes struct {
	Code    string `json:"code"`
	Status  string `json:"status"`
	Payload *[]TeamProfile
}

// team profile
type TeamProfile struct {
	Name             string   `json:"name"`
	ID               string   `json:"id"`
	Captain          string   `json:"captain"`
	Owner            string   `json:"owner"`
	Points           int`json:"points"`
	AccquiredMembers []db.TeamMembers `json:"accquiredMembers"`
	RelievedMembers  []db.TeamMembers `json:"relievedMembers"`
}

// User list resposne
type UserListRes struct {
	Code    string `json:"code"`
	Status  string `json:"status"`
	Payload *[]UserProfile
}


// User Profile
type UserProfile struct {
	Id          string `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Image       string  `json:"image"`
	AccessLevel string    `json:"accessLevel"`
	Points      int `json:"points,omitempty"`
	Status      string `json:"status"`
	UpdatedAt   int64 `json:"updatedAt"`
}

// Leader board response model
type UserLeaderBoard struct {
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	Email                   string `json:"email"`
	Level                   int    `json:"level"`
	PreviousLevelFinishTime string `json:"previous_level"`
}

type TickerRes struct {
	Payload struct {
				CurrentAuction struct {
								   ID        string `json:"Id"`
								   FirstName string `json:"FirstName"`
								   LastName  string `json:"LastName"`
								   Email     string `json:"Email"`
								   Image     string `json:"Image"`
								   Status    string `json:"Status"`
								   UpdatedAt int64    `json:"UpdatedAt"`
							   } `json:"current_auction"`
				TeamList       interface{} `json:"team_list"`
				AuctionHistory []AuctionHistoryRes `json:"auction_history"`
				TotalUsers int `json:"total_users"`
				AuctionedUsers int `json:"auctioned_users"`
			} `json:"payload"`
}

type AuctionHistoryRes struct {
	UserID string `json:"user_id"`
	UserName string `json:"user_name"`
	TeamID string `json:"team_id"`
	TeamName string `json:"team_name"`
	Action string `json:"action"`
	At     int64 `json:"at"`
}