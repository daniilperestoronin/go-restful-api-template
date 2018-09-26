package main

import (
	"log"
	"net/http"
)

const (
	dbDriver     = "postgres"
	dbDataSource = "user=postgres password=postgres dbname=records sslmode=disable"
)

func main() {

	recRep := NewRecordRepository(dbDriver, dbDataSource)
	recSer := NewRecordService(recRep)
	recHand := NewRecordHandler(recSer)

	http.HandleFunc("/record/", recHand.Handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
