package db

import "gopkg.in/mgo.v2/bson"

const UsersColl = "users"

// User schema for users collection
type User struct {
	Id          string `bson:"id"`
	FirstName   string    `bson:"firstName"`
	LastName    string    `bson:"lastName"`
	Email       string    `bson:"email"`
	Image       string  `bson:"image"`
	Password    string    `bson:"password"`
	AccessLevel string    `bson:"accessLevel"`
	AccessToken string    `bson:"accessToken"`
	Points      int `bson:"points"`
	Status      string `bson:"status"`
	UpdatedAt   int64 `bson:"updatedAt"`
}

func InsertNewUser(u *User) error {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(UsersColl)
	_, err := c.Upsert(bson.M{"id":u.Id}, &u)
	if err != nil {
		return err
	}
	return nil
}