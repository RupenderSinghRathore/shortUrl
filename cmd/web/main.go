package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":8080", "Port to be used for server")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := application{
		logger: logger,
	}

	logger.Info("Starting the server..", "port", *addr)
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
