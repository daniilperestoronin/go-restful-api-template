package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const connStr = "user=postgres password=postgres dbname=records sslmode=disable"

func createR(record Record) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec(
		"insert into record (title, text) values ( $1, $2)",
		record.title, record.text)
	if err != nil {
		panic(err)
	}
}
