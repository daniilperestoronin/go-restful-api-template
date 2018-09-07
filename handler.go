package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func recordHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		create(w, r)
	case http.MethodGet:
		read(w, r)
	case http.MethodPut:
		update(w, r)
	case http.MethodDelete:
		remove(w, r)
	default:
		bad(w, r)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	var record Record
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	createRecord(record)
}

func read(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == `/record/` {
		json.NewEncoder(w).Encode(readRecords())
	} else if regexp.MustCompile(`/record/+[0-9]+$`).MatchString(r.URL.Path) {
		id, err := strconv.ParseInt(
			strings.TrimPrefix(r.URL.Path, `/record/`),
			10,
			64)
		if err != nil {
			fmt.Fprintf(w, "Invalid id")
		}
		json.NewEncoder(w).Encode(readRecord(id))
	} else {
		bad(w, r)
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not yet implemented")
}

func remove(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not yet implemented")
}

func bad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
