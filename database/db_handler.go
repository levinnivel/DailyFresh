package database

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	var db *sql.DB
	var err error

	if db == nil && err == nil {
		db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/daily_fresh?parseTime=true&loc=Asia%2FJakarta")
		if err != nil {
			log.Fatal(err)
		}
	}

	return db
}
