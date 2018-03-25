package db

import (
	"log"

	"gopkg.in/mgo.v2"
)

const DB = "qbutsav_v1"

var Session *mgo.Session

func InitMongo() bool {
	url := "mongodb://localhost"
	log.Println("Establishing MongoDB connection...")
	var err error
	Session, err = mgo.Dial(url)
	if err != nil {
		log.Fatal("Cannot connect to MongoDB!")
		return true
	} else {
		log.Println("Connected to ", url)
		return false
	}
}

func GetSession() mgo.Session {
	return *Session.Copy()
}
