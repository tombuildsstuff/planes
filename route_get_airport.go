package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/tombuildsstuff/planes/models"
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

	flightsTo := make([]models.RouteSummary, 0)
	if v, ok := routes[code]; ok {
		for _, route := range v {
			flightsTo = append(flightsTo, models.RouteSummary{
				OriginCode:      code,
				DestinationCode: route.Destination,
				Uri:             fmt.Sprintf("/routes/%s/%s", code, route.Destination),
			})
		}
	}

	obj := models.GetAirportResponse{
		City:      airport.City,
		Code:      code,
		FullName:  airport.FullName,
		FlightsTo: flightsTo,
	}
	serializeJson(obj, w)
}
