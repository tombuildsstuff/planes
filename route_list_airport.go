package main

import (
	"fmt"
	"net/http"
)

// listAirports handles `/airports`
func listAirports(w http.ResponseWriter, r *http.Request) {
	obj := ListAirportsResponse{
		Airports: func() []Airport {
			out := make([]Airport, 0)

			for code, airport := range airports {
				out = append(out, Airport{
					Name: airport.FullName,
					Code: code,
					Uri: fmt.Sprintf("/airports/%s", code),
				})
			}

			return out
		}(),
	}
	serializeJson(obj, w)
}

type ListAirportsResponse struct {
	Airports []Airport `json:"airports"`
}
type Airport struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Uri string `json:"uri"`
}