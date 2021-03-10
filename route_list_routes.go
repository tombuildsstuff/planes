package main

import (
	"fmt"
	"net/http"
)

// listRoutes handles /routes
func listRoutes(w http.ResponseWriter, r *http.Request) {
	out := make([]RouteSummary, 0)
	for originCode, destinations := range routes {
		for _, destination := range destinations {
			out = append(out, RouteSummary{
				OriginCode:      originCode,
				DestinationCode: destination.Destination,
				Uri:             fmt.Sprintf("/route/%s/%s", originCode, destination.Destination),
			})
		}
	}
	obj := ListRoutesResponse{
		Routes: out,
	}
	serializeJson(obj, w)
}

type ListRoutesResponse struct {
	Routes []RouteSummary `json:"routes"`
}

type RouteSummary struct {
	OriginCode      string `json:"origin_code"`
	DestinationCode string `json:"destination_code"`
	Uri             string `json:"uri"`
}
