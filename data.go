package main

type AirportDetails struct {
	City     string
	FullName string
}

type RouteDetails struct {
	Destination         string
	TravelTimeInMinutes int32
}

func initialAirports() map[string]AirportDetails {
	return map[string]AirportDetails{
		"BER": {
			City:     "Berlin",
			FullName: "Berlin Brandenburg Airport",
		},
		"FCO": {
			City:     "Rome",
			FullName: "Roma Fiumicino Aeroporto",
		},
		"JFK": {
			City:     "New York",
			FullName: "John F. Kennedy International Airport",
		},
		"LHR": {
			City:     "London",
			FullName: "London Heathrow",
		},
		"SEA": {
			City:     "Seattle",
			FullName: "Seattle-Tacoma International Airport",
		},
	}
}

func initialRoutes() map[string][]RouteDetails {
	return map[string][]RouteDetails{
		"BER": {
			RouteDetails{
				Destination:         "FCO",
				TravelTimeInMinutes: 90,
			},
			RouteDetails{
				Destination:         "LHR",
				TravelTimeInMinutes: 60,
			},
		},
		"FCO": {
			RouteDetails{
				Destination:         "BER",
				TravelTimeInMinutes: 90,
			},
			RouteDetails{
				Destination:         "LHR",
				TravelTimeInMinutes: 180,
			},
		},
		"JFK": {
			RouteDetails{
				Destination:         "LHR",
				TravelTimeInMinutes: 460,
			},
			RouteDetails{
				Destination:         "SEA",
				TravelTimeInMinutes: 480,
			},
		},
		"LHR": {
			RouteDetails{
				Destination:         "BER",
				TravelTimeInMinutes: 60,
			},
			RouteDetails{
				Destination:         "JFK",
				TravelTimeInMinutes: 480,
			},
			RouteDetails{
				Destination:         "SEA",
				TravelTimeInMinutes: 660,
			},
		},
		"SEA": {
			RouteDetails{
				Destination:         "JFK",
				TravelTimeInMinutes: 480,
			},
			RouteDetails{
				Destination:         "LHR",
				TravelTimeInMinutes: 660,
			},
		},
	}
}
