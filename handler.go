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
		httpPost(w, r)
	case http.MethodGet:
		httpGet(w, r)
	case http.MethodPut:
		httpPut(w, r)
	case http.MethodDelete:
		httpDelete(w, r)
	default:
		bad(w, r)
	}
}

func httpPost(w http.ResponseWriter, r *http.Request) {
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
	create(record)
}

func httpGet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == `/record/` {
		records, err := readAll()
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		json.NewEncoder(w).Encode(records)
	} else if regexp.MustCompile(`/record/+[0-9]+$`).MatchString(r.URL.Path) {
		id, err := strconv.ParseInt(
			strings.TrimPrefix(r.URL.Path, `/record/`),
			10,
			64)
		if err != nil {
			fmt.Fprintf(w, "Invalid id")
		}
		record, err := read(id)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		json.NewEncoder(w).Encode(record)
	} else {
		bad(w, r)
	}
}

func httpPut(w http.ResponseWriter, r *http.Request) {
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
	update(record)
}

func httpDelete(w http.ResponseWriter, r *http.Request) {
	if regexp.MustCompile(`/record/+[0-9]+$`).MatchString(r.URL.Path) {
		id, err := strconv.ParseInt(
			strings.TrimPrefix(r.URL.Path, `/record/`),
			10,
			64)
		if err != nil {
			fmt.Fprintf(w, "Invalid id")
		}
		remove(id)
	} else {
		bad(w, r)
	}
}

func bad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
