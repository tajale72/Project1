package db

import (
	"database/sql"
	"log"
)

func PostgressClient() (*sql.DB, error) {
	conn := "host=localhost port=5432 user=romittajale dbname=finance sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected!")
	return db, nil
}
