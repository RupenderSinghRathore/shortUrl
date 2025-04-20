package main

import (
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl")
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Error(err.Error(), "Method", r.Method, "url", r.URL.RequestURI())
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
