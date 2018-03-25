package db

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const UsersColl = "users"

// User schema for users collection
type User struct {
	FirstName               string    `bson:"firstName"`
	LastName                string    `bson:"lastName"`
	Email                   string    `bson:"email"`
	Password                string    `bson:"password"`
	Level                   int       `bson:"level"`
	AccessLevel             string    `bson:"accessLevel"`
	AccessToken             string    `bson:"accessToken"`
	PreviousLevelFinishTime time.Time `bson:"previousLevelFinishTime"`
}

// Model for new user insert query
type InsertUserQuery struct {
	FirstName               string    `bson:"firstName"`
	LastName                string    `bson:"lastName"`
	Email                   string    `bson:"email"`
	Password                string    `bson:"password"`
	Level                   int       `bson:"level"`
	AccessLevel             string    `bson:"accessLevel"`
	AccessToken             string    `bson:"accessToken"`
	PreviousLevelFinishTime time.Time `bson:"previousLevelFinishTime"`
}

func GetUserByEmail(emailId string) (Team, error) {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(UsersColl)

	var user Team
	err := c.Find(bson.M{"email": emailId}).One(&user)
	if err != nil {
		return Team{}, err
	}
	return user, nil
}

func GetUserByAccessToken(t string) (Team, error) {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(UsersColl)

	var user Team
	err := c.Find(bson.M{"accessToken": t}).One(&user)
	if err != nil {
		return Team{}, err
	}
	return user, nil
}

func GetUserLeaderBoard(l int) (*[]Team, error) {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(UsersColl)

	var users []Team
	err := c.Find(bson.M{}).Sort("-level", "previousLevelFinishTime").Limit(l).All(&users)
	if err != nil {
		return &[]Team{}, err
	}
	return &users, nil
}

func InsertNewUser(u *InsertUserQuery) error {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(UsersColl)
	err := c.Insert(u)
	if err != nil {
		return err
	}
	return nil
}

func UpdateAccessTokenByEmailId(e string, t string) error {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(UsersColl)
	q := bson.M{"email": e}
	u := bson.M{"$set": bson.M{"accessToken": t}}
	err := c.Update(&q, &u)
	if err != nil {
		return err
	}
	return nil
}

func UpdateLevelByEmailId(e string, l int, t time.Time) error {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(UsersColl)
	q := bson.M{"email": e}
	u := bson.M{"$set": bson.M{"level": l, "previousLevelFinishTime": t}}
	err := c.Update(&q, &u)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePasswordByEmailId(e, p string) error {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(UsersColl)
	q := bson.M{"email": e}
	// Update password as well as log out current session
	u := bson.M{"$set": bson.M{"password": p, "accessToken": ""}}
	err := c.Update(&q, &u)
	if err != nil {
		return err
	}
	return nil
}

func UpdateRoleByEmailId(e string,r string) error {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(UsersColl)
	q := bson.M{"email": e}
	u := bson.M{"$set": bson.M{"accessLevel": r}}
	err := c.Update(&q, &u)
	if err != nil {
		return err
	}
	return nil
}
