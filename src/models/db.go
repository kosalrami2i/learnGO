package models

import (
	"log"

	"github.com/go-xorm/xorm"
)

/*Engine : This is the connection variable. */
var Engine *xorm.Engine

/*InitDB method establishes Database connection
 * and returns the same.
 */
func InitDB(driverName, connectionString string) {
	var err error
	Engine, err = xorm.NewEngine(driverName, connectionString)

	if err != nil {
		log.Fatal("Could not connect with database!!!!!!!!!!!!!!!!", err)
	}
}
