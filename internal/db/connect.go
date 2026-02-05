package db

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Must(ctx context.Context, url string) *sql.DB {
	if url == "" {
		log.Fatal("$DB_URL required")
	}

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("error connecting to db: ", err.Error())
	}

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("could not ping db: ", err.Error())
	}

	return db
}
