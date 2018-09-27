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

	recRep := NewPgRecordRepository(dbDriver, dbDataSource)
	recSer := NewRecordService(recRep)
	recHand := NewRecordHandler(recSer)

	http.HandleFunc("/record/", recHand.Handle)
	log.Println("Record service start")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
