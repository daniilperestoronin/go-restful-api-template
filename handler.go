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
	if r.URL.Path == `/record/` {
		json.NewEncoder(w).Encode(getRecords())
	} else if regexp.MustCompile(`/record/+[0-9]+$`).MatchString(r.URL.Path) {
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, `/record/`))
		if err != nil {
			fmt.Fprintf(w, "Invalid id")
		}
		json.NewEncoder(w).Encode(getRecord(id))
	} else {
		bad(w, r)
	}
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
