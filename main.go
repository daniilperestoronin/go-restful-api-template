package main

import (
	"log"
	"net/http"
	"os"

	"github.com/daniilperestoronin/go-restful-api-template/record"
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

	recRep := record.NewPgRepository(dbDriver, dbDataSource)
	recSer := record.NewService(recRep)
	recHand := record.NewHandler(recSer)

	http.HandleFunc("/record/", recHand.Handle)
	log.Println("Record service start")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
