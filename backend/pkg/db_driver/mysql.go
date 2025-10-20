package db_driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func GetInstance() *sql.DB {
	db_user := os.Getenv("DB_USER")
	//db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_database := os.Getenv("DB_DATABASE")
	db_port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s", db_user, db_host, db_port, db_database)
	fmt.Println(dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
