package main

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime/debug"

	"github.com/RupenderSinghRathore/shortUrl/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/pages/home.html")
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
	err := r.ParseForm()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	url := r.PostForm.Get("url")
	if url == "" {
		http.Error(w, "$$ Url field is empty!! $$", http.StatusBadRequest)
		return
	}
	hash := models.CreateHash(url)
	app.urlDB.Insert(url, hash)
	w.Write([]byte(hash))
}

func (app *application) redirect(w http.ResponseWriter, r *http.Request) {
	hash := r.PathValue("id")
	url, err := app.urlDB.Retrieve(hash)
	// print("url:", url)
	if err != nil {
		println("url if error:", url)
		app.serverError(w, r, err)
		return
	}
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Error(err.Error(), "Method", r.Method, "url", r.URL.RequestURI())
	fmt.Println(string(debug.Stack()))
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
