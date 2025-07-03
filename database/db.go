package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDb() *sql.DB {
	dns := "host=localhost user=postgres password=postgres dbname=system_inventory sslmode=disable"
	db, err := sql.Open("postgres", dns)
	if err != nil {
		log.Fatal(err)
	}
	
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	DB = db
	return  db
}


