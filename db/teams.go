package db

import "gopkg.in/mgo.v2/bson"

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
	Id        string `bson:"id"`
	FirstName string    `bson:"firstName"`
	LastName  string    `bson:"lastName"`
	Email     string    `bson:"email"`
	Image     string  `bson:"image"`
	Point     int `bson:"point"`
	UpdatedAt int64 `bson:"updatedAt"`
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

func GetTeamById(id string) (*Team, error) {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(TeamsColl)

	var teams Team
	err := c.Find(bson.M{"id":id}).One(&teams)
	if err != nil {
		return &Team{}, err
	}
	return &teams, nil
}

func UpdateTeamById(t *Team) error {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(TeamsColl)

	q := bson.M{"id": t.ID}
	err := c.Update(&q, t)
	if err != nil {
		return err
	}
	return nil
}

func DeallocateUserById(t *Team, u *User) error {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(TeamsColl)

	q := bson.M{"id": t.ID}
	d := bson.M{
		"$pull":bson.M{"accquiredMembers":bson.M{"id":u.Id}},
		"$set":bson.M{"points":t.Points},
	}
	err := c.Update(&q, d)
	if err != nil {
		return err
	}
	return nil
}