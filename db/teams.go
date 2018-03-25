package db

const TeamsColl = "teams"

// Team schema for users collection
type Team struct {
	Name             string   `bson:"name"`
	ID               string   `bson:"id"`
	Captain          string   `bson:"captain"`
	Owner            string   `bson:"owner"`
	Points           int    `bson:"points"`
	AccquiredMembers []string `bson:"accquiredMembers"`
	RelievedMembers  []string `bson:"relievedMembers"`
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