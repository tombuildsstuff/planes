package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var airports map[string]AirportDetails
var routes map[string][]RouteDetails

func main() {
	airports  = initialAirports()
	routes = initialRoutes()

	router := mux.NewRouter()
	registerRoutes(router)
	http.Handle("/", router)

	server := &http.Server{
		Handler: router,
		Addr:    "localhost:2021",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Check out: http://%s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("error: %+v", err)
		os.Exit(1)
		return
	}

	os.Exit(0)
}
