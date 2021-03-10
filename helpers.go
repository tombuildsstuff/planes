package main

import (
	"encoding/json"
	"net/http"
)

func serializeJson(obj interface{}, w http.ResponseWriter) {
	if err := json.NewEncoder(w).Encode(&obj); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
