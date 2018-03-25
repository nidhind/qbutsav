// Types models in HTTP Requests

package models

// For new team registeration
type NewTeam struct {
	Name             string   `json:"name"`
	ID               string   `json:"id"`
	Captain          string   `json:"captain"`
	Owner            string   `json:"owner"`
	Points           int      `json:"points"`
	AccquiredMembers []string `json:"accquiredMembers"`
	RelievedMembers  []string `json:"relievedMembers"`
}

// For accesstoken request
type GenAccessToken struct {
	Email    string `json:"emailId"`
	Password string `json:"password"`
}

//For adding new puzzles
type PuzzleReq struct {
	Level        int    `json:"level,string"`
	Image        string `json:"image"`
	Clue         string `json:"clue"`
	SolutionHash string `json:"solutionHash"`
}

//For submitting answers
type AnswerReq struct {
	Answer string `json:"answer"`
}

// For forgot password email request
type ResetPswdEmailReq struct {
	Email string `json:"email_id"`
}

// For forgot password update request
type ResetPswdUpdateReq struct {
	NewPassword string `json:"new_password"`
}


//For updating user role
type RoleUpdateReq struct {
	Email string `json:"emailId"`
	AccessLevel string `json:"accessLevel"`
}
