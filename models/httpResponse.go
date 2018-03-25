package models

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
	AccquiredMembers []string `json:"accquiredMembers"`
	RelievedMembers  []string `json:"relievedMembers"`
}

// Leader board response model
type UserLeaderBoard struct {
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	Email                   string `json:"email"`
	Level                   int    `json:"level"`
	PreviousLevelFinishTime string `json:"previous_level"`
}
