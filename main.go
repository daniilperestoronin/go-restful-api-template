package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	var (
		dbDriver     = os.Getenv("DB_DRIVER")
		dbDataSource = os.Getenv("DB_DATA_SOURCE")
	)
	if dbDriver == "" || dbDataSource == "" {
		panic("DB_DRIVER and DB_DATA_SOURCE not specified")
	}
	log.Print("DB_DRIVER: " + dbDriver)
	log.Print("DB_DATA_SOURCE: " + dbDataSource)

	recRep := NewPgRecordRepository(dbDriver, dbDataSource)
	recSer := NewRecordService(recRep)
	recHand := NewRecordHandler(recSer)

	http.HandleFunc("/record/", recHand.Handle)
	log.Println("Record service start")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
