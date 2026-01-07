package handler

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	ts := template.New("home")

	files := []string{
		"pkg/admin/api/handler/templates/home.html",
	}

	ts, err := ts.ParseFiles(files...)
	if err != nil {
		slog.Info(fmt.Sprintf("ParseFiles error: %s", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	err = ts.ExecuteTemplate(w, "home", nil)
	if err != nil {
		slog.Info(fmt.Sprintf("Execute error: %s", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
// ejecuta el template por el nombre dado
