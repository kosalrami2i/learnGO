package models

import (
	"database/sql"
	"log"
	"reflect"

	_ "github.com/lib/pq"
)

func InitDB(driverName, connectionString string) {
	db, err := sql.Open(driverName, connectionString)

	if err != nil {
		log.Fatal("Could not connect with database!!!!!!!!!!!!!!!!", err)
	}

	log.Print(reflect.TypeOf(db))
}
