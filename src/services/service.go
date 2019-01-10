package services

import (
	"encoding/json"
	"fmt"
	"models"
	"reflect"
)

/*AddUser function gets the JSON Object and adds in the database. */
func AddUser(jsonData []uint8) {
	var user models.Users

	json.Unmarshal(jsonData, &user)
	fmt.Println("Added User ::::::::::::::::::::::::::::", reflect.TypeOf(models.Engine))
	fmt.Println(user)
	affected, err := models.Engine.Insert(user)

	if err != nil {
		fmt.Println("Error Adding User", err)
		return
	}
	fmt.Println("::::::::::::::::Affecteddddddd::::::::::::::::::::", affected)

}
