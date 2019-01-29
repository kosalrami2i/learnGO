package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"models"
	"services"

	_ "github.com/lib/pq"
)

func main() {
	// connectionString := "postgres://postgres:ubuntu@192.168.2.45:5432/learnGO"
	connectionString := "host=192.168.2.45 port=5432 user=postgres password=ubuntu dbname=learnGO sslmode=disable"
	driverName := "postgres"
	models.InitDB(driverName, connectionString)

	fmt.Println("Happy Learning.....!!!!!!")

	jsonFile, err := os.Open("src/data.json")
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

	services.AddUser(jsonData)
}
