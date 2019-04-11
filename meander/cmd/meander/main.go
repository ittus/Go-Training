package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/ittus/Go-Training/meander"
)

func main() {
	meander.APIKey = os.Getenv("GOOGLE_PLACE_APIKEY")
	http.HandleFunc("/journeys", func(w http.ResponseWriter,
		r *http.Request) {
		respond(w, r, meander.Journeys)
	})
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	publicData := make([]interface{}, len(data))
	for i, d := range data {
		publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}
