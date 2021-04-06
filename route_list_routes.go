package main

import (
	"fmt"
	"net/http"

	"github.com/tombuildsstuff/planes/models"
)

// listRoutes handles /routes
func listRoutes(w http.ResponseWriter, r *http.Request) {
	out := make([]models.RouteSummary, 0)
	for originCode, destinations := range routes {
		for _, destination := range destinations {
			out = append(out, models.RouteSummary{
				OriginCode:      originCode,
				DestinationCode: destination.Destination,
				Uri:             fmt.Sprintf("/route/%s/%s", originCode, destination.Destination),
			})
		}
	}
	obj := models.ListRoutesResponse{
		Routes: out,
	}
	serializeJson(obj, w)
}
