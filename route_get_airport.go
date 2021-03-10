package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// getAirport handles /airports/{code}
func getAirport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	airport, ok := airports[code]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	flightsTo := make([]RouteSummary, 0)
	if v, ok := routes[code]; ok {
		for _, route := range v {
			flightsTo = append(flightsTo, RouteSummary{
				OriginCode: code,
				DestinationCode: route.Destination,
				Uri: fmt.Sprintf("/routes/%s/%s", code, route.Destination),
			})
		}
	}

	obj := GetAirportResponse{
		City:     airport.City,
		Code:     code,
		FullName: airport.FullName,
		FlightsTo: flightsTo,
	}
	serializeJson(obj, w)
}

type GetAirportResponse struct {
	City string `json:"city"`
	Code string `json:"code"`
	FullName string `json:"full_name"`
	FlightsTo []RouteSummary `json:"flights_to"`
}