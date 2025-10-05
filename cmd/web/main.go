package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/RupenderSinghRathore/shortUrl/internal/models"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type application struct {
	logger *slog.Logger
	urlDB  *models.UrlDB
}

func main() {
	addr := flag.String("addr", "0.0.0.0:8080", "Port to be used for server")
	dsn := flag.String(
		"dsn",
		"postgres://kami-sama:touka@localhost:5432/shorturl",
		"Postgres connection string",
	)
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()
	app := application{
		logger: logger,
		urlDB:  &models.UrlDB{DB: db},
	}

	logger.Info("Starting the server..", "port", *addr)
	// err = http.ListenAndServeTLS(*addr, "server.crt", "server.key", app.routes())
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		defer db.Close()
		return nil, err
	}
	return db, nil
}
