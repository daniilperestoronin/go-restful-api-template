package main

import (
	"encoding/json"
	"net/http"
)

func recordHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(getRecord())
	case http.MethodPost:
		// Create a new record.
	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		w.WriteHeader(500)
	}
}

func getRecord() Record {
	return Record{Id: 1, Title: "Title", Text: "Blah Blah Blah Blah"}
}
