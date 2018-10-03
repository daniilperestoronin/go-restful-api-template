package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	var dbDriver, dbDataSource string
	flag.StringVar(&dbDriver, "db_driver", "bar", "a string var")
	flag.StringVar(&dbDataSource, "db_data_source", "bar", "a string var")
	flag.Parse()

	if dbDriver == "" || dbDataSource == "" {
		panic("db_driver or b_data_source don't specified")
	}

	recRep := NewPgRecordRepository(dbDriver, dbDataSource)
	recSer := NewRecordService(recRep)
	recHand := NewRecordHandler(recSer)

	http.HandleFunc("/record/", recHand.Handle)
	log.Println("Record service start")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
