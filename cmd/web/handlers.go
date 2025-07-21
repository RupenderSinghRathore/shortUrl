package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/RupenderSinghRathore/shortUrl/internal/models"
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

func (app *application) giveHash(w http.ResponseWriter, r *http.Request) {
	var url = &models.UrlStr{}
	err := json.NewDecoder(r.Body).Decode(url)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	if url.Url == "" {
		http.Error(w, "$$ Url field is empty!! $$", http.StatusBadRequest)
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

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
