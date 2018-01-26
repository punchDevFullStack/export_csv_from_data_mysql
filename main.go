package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// var data = [][]string{{"Line1", "Hello Readers of", "Line2", "golangcode.com"}}

func main() {
	db, err := gorm.Open("mysql", "root:1234@(localhost:3306)/testGolang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var members []Member
	db.Find(&members)

	file, err := os.Create("result.csv")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headColumn := []string{"first_name", "last_name", "email", "password"}
	errHead := writer.Write(headColumn)
	if errHead != nil {
		log.Fatal("Cannot write to file", errHead)
	}

	data := mapMemberString(members)
	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}
}

type Member struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `gorm:"primary_key" json:"email"`
	Password  string `json:"password"`
}

func mapMemberString(members []Member) [][]string {
	var ret [][]string

	for _, value := range members {
		row := []string{
			value.FirstName,
			value.LastName,
			value.Email,
			value.Password,
		}

		ret = append(ret, row)

	}
	fmt.Println("total value ", ret)
	return ret
}
