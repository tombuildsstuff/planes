package main

import (
	"fmt"
	"net/http"

	"github.com/tombuildsstuff/planes/models"
)

// listAirports handles `/airports`
func listAirports(w http.ResponseWriter, r *http.Request) {
	obj := models.ListAirportsResponse{
		Airports: func() []models.Airport {
			out := make([]models.Airport, 0)

			for code, airport := range airports {
				out = append(out, models.Airport{
					Name: airport.FullName,
					Code: code,
					Uri:  fmt.Sprintf("/airports/%s", code),
				})
			}

			return out
		}(),
	}
	serializeJson(obj, w)
}
