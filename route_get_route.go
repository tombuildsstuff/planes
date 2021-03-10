package main

import (
	"github.com/gorilla/mux"
	"net/http"
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


	obj := GetRouteResponse{
		OriginAirportCode:      originCode,
		OriginAirportName:      fromAirport.FullName,
		DestinationAirportCode: destinationCode,
		DestinationAirportName: toAirport.FullName,
		FlightLengthInMinutes:  route.TravelTimeInMinutes,
	}
	serializeJson(obj, w)
}

type GetRouteResponse struct {
	OriginAirportCode string `json:"origin_code"`
	OriginAirportName string `json:"origin_name"`

	DestinationAirportCode string `json:"destination_code"`
	DestinationAirportName string `json:"destination_name"`

	FlightLengthInMinutes int32 `json:"flight_length_in_minutes"`
}