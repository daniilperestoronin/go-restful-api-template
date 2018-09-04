package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var records map[int]Record

func recordHandler(w http.ResponseWriter, r *http.Request) {
	records = make(map[int]Record)
	records[1] = Record{Id: 1, Title: "Title 1", Text: "Blah Blah Blah Blah 1"}
	records[2] = Record{Id: 2, Title: "Title 2", Text: "Blah Blah Blah Blah 2"}
	records[3] = Record{Id: 3, Title: "Title 3", Text: "Blah Blah Blah Blah 3"}

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

func getRecord() map[int]Record {
	return records
}
