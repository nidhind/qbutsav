package main

import (
	"os"
	"bufio"
	"io"
	"log"
	"fmt"
	"encoding/csv"
	"github.com/nidhind/qbutsav/db"
)

func main() {
	if len(os.Args) !=3{
		log.Fatalln("Please provide filepath as argument 1 " +
			"and profile image base path as argument 2")
	}
	filePath:=os.Args[1]
	PROFILE_IMG_BASE_URL:=os.Args[2]
	db.InitMongo()
	csvFile, _ := os.Open(filePath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	c := 0
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Panicln(err)
		}
		c++
		fmt.Println(c)

		u := db.User{
			Id:line[0],
			FirstName:line[2],
			LastName: line[3],
			Email:line[1],
			Image: fmt.Sprintf("%s/%s.jpg", PROFILE_IMG_BASE_URL,line[0]),
			Status:"waiting",
		}
		err = db.InsertNewUser(&u)
		if err != nil {
			fmt.Println("Error at ", c, err)
		}
	}
	fmt.Println("Completed ", c)
}