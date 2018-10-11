package record

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Handler struct {
	srv Service
}

func NewHandler(s Service) Handler {
	return Handler{
		srv: s,
	}
}

func (h Handler) Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("Request from: " + r.RemoteAddr +
		", method:" + r.Method +
		", to: " + r.URL.Host)
	switch r.Method {
	case http.MethodPost:
		h.httpPost(w, r)
	case http.MethodGet:
		h.httpGet(w, r)
	case http.MethodPut:
		h.httpPut(w, r)
	case http.MethodDelete:
		h.httpDelete(w, r)
	default:
		h.bad(w, r)
	}
}

func (h Handler) httpPost(w http.ResponseWriter, r *http.Request) {
	var record Record
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		log.Print("Missing request body, return 400")
		return
	}
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		http.Error(w, err.Error(), 400)
		log.Print("Wrong request body, return 400: ", err.Error())
		return
	}
	h.srv.create(record)
}

func (h Handler) httpGet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == `/record/` {
		records, err := h.srv.readAll()
		if err != nil {
			http.Error(w, err.Error(), 500)
			log.Print("Get Records error, return 500: ", err.Error())
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
		record, err := h.srv.read(id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			log.Print("Get Record error, return 500: ", err.Error())
		}
		json.NewEncoder(w).Encode(record)
	} else {
		h.bad(w, r)
	}
}

func (h Handler) httpPut(w http.ResponseWriter, r *http.Request) {
	var record Record
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		log.Print("Wrong request body, return 400")
		return
	}
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		http.Error(w, err.Error(), 400)
		log.Print("update Error, return 400")
		return
	}
	h.srv.update(record)
}

func (h Handler) httpDelete(w http.ResponseWriter, r *http.Request) {
	if regexp.MustCompile(`/record/+[0-9]+$`).MatchString(r.URL.Path) {
		id, err := strconv.ParseInt(
			strings.TrimPrefix(r.URL.Path, `/record/`),
			10,
			64)
		if err != nil {
			fmt.Fprintf(w, "Invalid id")
		}
		h.srv.remove(id)
	} else {
		h.bad(w, r)
	}
}

func (h Handler) bad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
