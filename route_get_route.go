package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/tombuildsstuff/planes/models"
)

// getRoute handles /routes/{origin}/{destination}
func getRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	originCode := vars["origin"]
	destinationCode := vars["destination"]
	if originCode == "" || destinationCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fromAirport, ok := airports[originCode]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	toAirport, ok := airports[destinationCode]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	flightsFromOrigin, ok := routes[originCode]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var route *RouteDetails
	for _, flight := range flightsFromOrigin {
		if flight.Destination == destinationCode {
			route = &flight
			break
		}
	}
	if route == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	obj := models.GetRouteResponse{
		OriginAirportCode:      originCode,
		OriginAirportName:      fromAirport.FullName,
		DestinationAirportCode: destinationCode,
		DestinationAirportName: toAirport.FullName,
		FlightLengthInMinutes:  route.TravelTimeInMinutes,
	}
	serializeJson(obj, w)
}
