package main

import (
	"github.com/gorilla/mux"
)

func registerRoutes(router *mux.Router) {
	router.HandleFunc("/", information)
	router.HandleFunc("/airports", listAirports)
	router.HandleFunc("/airports/{code}", getAirport)
	router.HandleFunc("/routes", listRoutes)
	router.HandleFunc("/routes/{origin}/{destination}", getRoute)
}
