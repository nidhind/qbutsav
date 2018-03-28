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

	db.InitMongo()
	filePath:=os.Args[1]
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
			Status:"waiting",
		}
		err = db.InsertNewUser(&u)
		if err != nil {
			fmt.Println("Error at ", c, err)
		}
	}
	fmt.Println("Completed ", c)
}