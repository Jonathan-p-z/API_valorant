package controllers

import (
	"html/template"
	"net/http"
)

func ErrorHandler(status int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		tmpl, err := template.ParseFiles(
			"templates/header.html",
			"templates/footer.html",
			"templates/error.html",
		)
		if err != nil {
			http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.ExecuteTemplate(w, "error.html", map[string]interface{}{
			"Status": status,
		})
	}
}
