package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	slog.Info(fmt.Sprintf("Starting web server at localhost:%s\n", "8081"))

	if err := http.ListenAndServe(":8000", NewRouter()); err != nil {
		slog.Error("Could not start webserver", "err_msg", err.Error())
	}
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHello)

	return mux
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	slog.Info(fmt.Sprintf("%s %s\n", r.Method, r.URL.Path))

	_, err := fmt.Fprintf(w, "%s\n", hello())
	if err != nil {
		slog.Error(err.Error())
	}
}

func hello() string {
	return "Hello World"
}
