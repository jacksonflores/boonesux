package db

import (
	"database/sql"
	"log"
	"os"
)

func Connect() {
	url := os.Getenv("DB_URL")
	if url == "" {
		log.Fatal("$DB_URL required")
	}

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("error connecting to db: ", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("could not ping db: ", err.Error())
	}
}
