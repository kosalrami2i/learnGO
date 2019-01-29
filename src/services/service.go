package services

import (
	"encoding/json"
	"fmt"
	"models"
)

/*AddUser function gets the JSON Object and adds in the database. */
func AddUser(jsonData []uint8) {
	user := new(models.Users)

	json.Unmarshal(jsonData, user)
	fmt.Println("User================================>>>>>>>>>>>>>>>>", string(jsonData))

	models.Db.NewRecord(user)
	affected := models.Db.Create(&user)

	fmt.Println("User Created :::::::::::::::::::::::", models.Db.NewRecord(user), affected)
}
