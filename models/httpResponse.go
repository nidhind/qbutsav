package models

type LoginRes struct {
	Code    string `json:"code"`
	Status  string `json:"status"`
	Payload struct {
		AccessToken string `json:"access_token"`
	}
}

type ProfileRes struct {
	Code    string `json:"code"`
	Status  string `json:"status"`
	Payload interface{}
}

// User schema for users collection
type UserProfile struct {
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	Email                   string `json:"email"`
	Level                   int    `json:"level"`
	LevelImage              string `json:"level_image"`
	LevelClue               string `json:"level_clue"`
	AccessLevel             string ` json:"access_level"`
	PreviousLevelFinishTime string `json:"previous_level"`
}

// Leader board response model
type UserLeaderBoard struct {
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	Email                   string `json:"email"`
	Level                   int    `json:"level"`
	PreviousLevelFinishTime string `json:"previous_level"`
}
