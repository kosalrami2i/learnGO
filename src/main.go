package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"src/models"

	_ "github.com/lib/pq"
)

/*Address is a derived datatype that gets address data of a user and constructs a JSON object */
type Address struct {
	ID       int    `json:"id"`
	Location string `json:"location"`
}

/*User is a derived datatype that gets user data and constructs a JSON object */
type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	MobileNumber string    `json:"mobile_number"`
	Email        string    `json:"email"`
	Addresses    []Address `json:"addresses"`
}

func main() {
	connectionString := "user=postgres dbname=learnGO  password=ubuntu"
	driverName := "postgres"
	models.InitDB(driverName, connectionString)
	fmt.Println("Happy Learning.....!!!!!!")

	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error reading JSON file", err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data", err)
		return
	}
	fmt.Println(string(jsonData))

	var user User
	json.Unmarshal(jsonData, &user)
	fmt.Println(user)
}
