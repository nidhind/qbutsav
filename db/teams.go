package db

const TeamsColl = "teams"

// Team schema for users collection
type Team struct {
	Name             string   `bson:"name"`
	ID               string   `bson:"id"`
	Captain          string   `bson:"captain"`
	Owner            string   `bson:"owner"`
	AccquiredMembers []string `bson:"accquiredMembers"`
	RelievedMembers  []string `bson:"relievedMembers"`
}

func GetAllTeams() ([]Team, error) {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(TeamsColl)

	var teams []Team
	err := c.Find(nil).All(&teams)
	if err != nil {
		return []Team{}, err
	}
	return teams, nil
}
