package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	// handlers for API
	getR := sm.Methods(http.MethodPost).Subrouter()
	getR.HandleFunc("/mail/{sender:[A-z0-9]+}/{recieve:[A-z0-9]+}/{subject:[A-z0-9]+}/{body:[A-z0-9]+}", GetThatMail)

	// create a new server
	s := http.Server{
		Addr:         ":9095",           // configure the bind address
		Handler:      sm,                // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
	// start the server
  fmt.Println("Starting server on port 9095")
}
