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
