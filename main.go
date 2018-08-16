package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Record struct {
	Id    int64  `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
}

func recordHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		record := Record{Id: 1, Title: "Title", Text: "Blah Blah Blah Blah"}
		json.NewEncoder(w).Encode(record)
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

func main() {
	http.HandleFunc("/record/", recordHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
