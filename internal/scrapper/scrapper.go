package scrapper

import (
	"encoding/json"
	"net/http"
)

type ScrapeRequest struct {
	product string
}

func (c *Controller) Scrape() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ScrapeRequest
		json.NewDecoder(r.Body).Decode(&req)

		msg := "Hello world"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&msg)
	}
}
