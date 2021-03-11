package models

type RouteSummary struct {
	OriginCode      string `json:"origin_code"`
	DestinationCode string `json:"destination_code"`
	Uri             string `json:"uri"`
}

type Airport struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Uri  string `json:"uri"`
}

type ListAirportsResponse struct {
	Airports []Airport `json:"airports"`
}

type GetAirportResponse struct {
	City      string         `json:"city"`
	Code      string         `json:"code"`
	FullName  string         `json:"full_name"`
	FlightsTo []RouteSummary `json:"flights_to"`
}

type ListRoutesResponse struct {
	Routes []RouteSummary `json:"routes"`
}

type GetRouteResponse struct {
	OriginAirportCode string `json:"origin_code"`
	OriginAirportName string `json:"origin_name"`

	DestinationAirportCode string `json:"destination_code"`
	DestinationAirportName string `json:"destination_name"`

	FlightLengthInMinutes int32 `json:"flight_length_in_minutes"`
}
