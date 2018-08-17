package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/record/", recordHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
