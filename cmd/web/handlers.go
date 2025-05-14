package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/RupenderSinghRathore/shortUrl/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl")
	if err != nil {
		log.Fatalf("Error: %v   :in fuck function", err)
	}
	err = ts.Execute(w, nil)
	// files := []string{
	// 	"./ui/html/pages/home.tmpl",
	// 	"./ui/html/base.tmpl",
	// }
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }
	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }
}

func (app *application) giveHash(w http.ResponseWriter, r *http.Request) {
	var url = &models.UrlStr{}
	err := json.NewDecoder(r.Body).Decode(url)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	hash := models.CreateHash(url.Url)
	app.urlDB.Insert(url.Url, hash)
	w.Write([]byte(hash))
}

func (app *application) redirect(w http.ResponseWriter, r *http.Request) {
	hash := r.PathValue("id")
	url, err := app.urlDB.Retrieve(hash)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Error(err.Error(), "Method", r.Method, "url", r.URL.RequestURI())
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
