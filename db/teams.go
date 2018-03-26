package db

const TeamsColl = "teams"

// Team schema for users collection
type Team struct {
	Name             string   `bson:"name"`
	ID               string   `bson:"id"`
	Captain          string   `bson:"captain"`
	Owner            string   `bson:"owner"`
	Points           int    `bson:"points"`
	AccquiredMembers []TeamMembers `bson:"accquiredMembers"`
	RelievedMembers  []TeamMembers `bson:"relievedMembers"`
}

type TeamMembers struct {
	Id          string `bson:"id"`
	FirstName   string    `bson:"firstName"`
	LastName    string    `bson:"lastName"`
	Email       string    `bson:"email"`
	Image       string  `bson:"image"`
	UpdatedAt   int64 `bson:"updatedAt"`
}

func GetAllTeams() (*[]Team, error) {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(TeamsColl)

	var teams []Team
	err := c.Find(nil).All(&teams)
	if err != nil {
		return &[]Team{}, err
	}
	return &teams, nil
}

func CreateNewTeam(t *Team) error {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(TeamsColl)
	err := c.Insert(t)
	if err != nil {
		return err
	}
	return nil
}