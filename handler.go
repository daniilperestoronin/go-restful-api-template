package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func recordHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		get(w, r)
	case http.MethodPost:
		post(w, r)
	case http.MethodPut:
		put(w, r)
	case http.MethodDelete:
		delete(w, r)
	default:
		bad(w, r)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getRecord())
}

func post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not yet implemented")
}

func put(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not yet implemented")
}

func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not yet implemented")
}

func bad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func getRecord() Record {
	return Record{Id: 1, Title: "Title", Text: "Blah Blah Blah Blah"}
}
