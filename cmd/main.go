package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jacksonflores/boonesux/internal/app"
	"github.com/jacksonflores/boonesux/internal/db"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("$DB_URL required")
	}

	db := db.Must(ctx, dbUrl)
	defer db.Close()

	app := app.New(db)

	if err := app.Run(ctx, "8080"); err != nil {
		log.Fatal(err)
	}
}
