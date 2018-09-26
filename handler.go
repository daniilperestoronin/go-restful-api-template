package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type RecordHandler struct {
	rSrv RecordService
}

func NewRecordHandler(recSrv RecordService) RecordHandler {
	return RecordHandler{
		rSrv: recSrv,
	}
}

func (rh RecordHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		rh.httpPost(w, r)
	case http.MethodGet:
		rh.httpGet(w, r)
	case http.MethodPut:
		rh.httpPut(w, r)
	case http.MethodDelete:
		rh.httpDelete(w, r)
	default:
		rh.bad(w, r)
	}
}

func (rh RecordHandler) httpPost(w http.ResponseWriter, r *http.Request) {
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
	rh.rSrv.Create(record)
}

func (rh RecordHandler) httpGet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == `/record/` {
		records, err := rh.rSrv.ReadAll()
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
		record, err := rh.rSrv.Read(id)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		json.NewEncoder(w).Encode(record)
	} else {
		rh.bad(w, r)
	}
}

func (rh RecordHandler) httpPut(w http.ResponseWriter, r *http.Request) {
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
	rh.rSrv.Update(record)
}

func (rh RecordHandler) httpDelete(w http.ResponseWriter, r *http.Request) {
	if regexp.MustCompile(`/record/+[0-9]+$`).MatchString(r.URL.Path) {
		id, err := strconv.ParseInt(
			strings.TrimPrefix(r.URL.Path, `/record/`),
			10,
			64)
		if err != nil {
			fmt.Fprintf(w, "Invalid id")
		}
		rh.rSrv.Remove(id)
	} else {
		rh.bad(w, r)
	}
}

func (rh RecordHandler) bad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
