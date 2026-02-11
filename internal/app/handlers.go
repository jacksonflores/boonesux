package app

import "net/http"

func (a *App) healthz(w http.ResponseWriter, r *http.Request) {
	err := a.Templates.ExecuteTemplate(w, "health.html", nil)
	if err != nil {
		http.Error(w, "err", http.StatusServiceUnavailable)
	}
}

func (a *App) about(w http.ResponseWriter, r *http.Request) {
	err := a.Templates.ExecuteTemplate(w, "ramblings.html", nil)
	if err != nil {
		http.Error(w, "err", http.StatusInternalServerError)
	}
}
