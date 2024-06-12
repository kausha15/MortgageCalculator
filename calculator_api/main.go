package main

import (
	"fmt"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"quoter_assignment/calculator_api/handlers"
	"time"
)

// getEnvOrDefault gets environment variable matching key.
// If key not found returns supplied default
func getEnvOrDefault(k string, def string) string {
	v := os.Getenv(k)
	if v == "" {
		return def
	}

	return v
}

func main() {
	p := getEnvOrDefault("HTTP_PORT", "8080")
	l := log.New(os.Stdout, "calculate-api ", log.LstdFlags)

	h := handlers.NewCalculator(l)
	r := mux.NewRouter()

	// POST /calculate
	r.HandleFunc("/calculate", h.Calculate).Methods(http.MethodPost)

	// public api
	c := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}), gohandlers.AllowedHeaders([]string{"content-type"}))

	l.Printf("starting server on port %s", p)
	s := http.Server{
		Addr:         fmt.Sprintf(":%s", p),
		Handler:      c(r),
		ErrorLog:     l,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	l.Fatal(s.ListenAndServe())
}
