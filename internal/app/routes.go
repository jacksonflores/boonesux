package app

import "net/http"

func (a *App) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", a.healthz)
	mux.HandleFunc("/", a.about)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))

	return mux
}
