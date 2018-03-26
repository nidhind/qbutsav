package db

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

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

func GetAllUsers() (*[]User, error) {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(UsersColl)

	var users []User
	err := c.Find(nil).All(&users)
	if err != nil {
		return &[]User{}, err
	}
	return &users, nil
}

func GetUserById(id string) (User, error) {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(UsersColl)

	var user User
	err := c.Find(bson.M{"id": id}).One(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func GetUsersByStatus(st string) ([]User, error) {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(UsersColl)

	var users []User
	err := c.Find(bson.M{"status": st}).All(&users)
	if err != nil {
		return []User{}, err
	}
	return users, nil
}


func UpdateUserStatusById(id string, st string) error {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(UsersColl)

	q := bson.M{"id": id}
	u := bson.M{"$set": bson.M{"status": st}}
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

func UpdateRoleByEmailId(e string, r string) error {
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
