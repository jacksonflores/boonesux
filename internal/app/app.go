package app

import (
	"context"
	"database/sql"
	"errors"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

type App struct {
	DB        *sql.DB
	Server    *http.Server
	Templates *template.Template
}

func New(db *sql.DB) *App {
	pattern := filepath.Join("web", "templates", "*.html")
	templates, err := template.ParseGlob(pattern)
	if err != nil {
		log.Fatalf("error parsing templates: %v", err)
	}
	return &App{
		DB:        db,
		Templates: templates,
	}
}

func (app *App) Run(ctx context.Context, port string) error {
	app.Server = &http.Server{
		Addr:              ":" + port,
		Handler:           app.Routes(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		log.Printf("listening on %s", app.Server.Addr)
		if err := app.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		log.Println("shutdown requested")
	case err := <-errCh:
		return err

	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return app.Server.Shutdown(shutdownCtx)
}
