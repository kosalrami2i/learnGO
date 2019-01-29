package models

import (
	"fmt"
	"log"
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

/*InitDB method establishes Database connection
 * and returns the same.
 */
func InitDB(driverName, connectionString string) {
	var err error

	Db, err = gorm.Open(driverName, connectionString)
	fmt.Println(reflect.TypeOf(Db))

	if err != nil {
		log.Fatal("Could not connect with database!!!!!!!!!!!!!!!!", err)
	}
}
