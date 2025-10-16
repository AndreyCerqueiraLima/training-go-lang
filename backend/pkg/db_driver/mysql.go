package db_driver

import (
	"database/sql"
	"log"
)

func GetInstance() *sql.DB {
	dsn := "root:@tcp(localhost:3306)/training_test_go"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
