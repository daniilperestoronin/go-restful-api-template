package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const connStr = "user=postgres password=postgres dbname=records sslmode=disable"

type RecordRepository struct{}

func NewRecordRepository() RecordRepository {
	return RecordRepository{}
}

func (r *RecordRepository) Create(record Record) (int64, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var id, errIn = db.Exec(
		`INSERT INTO record (title, text) VALUES ( $1, $2)`,
		record.Title, record.Text)
	if errIn != nil {
		return 0, err
	}
	return id.RowsAffected()
}

func (r *RecordRepository) ReadAll() ([]Record, error) {
	db, err := sql.Open("postgres", connStr)
	rows, err := db.Query(`SELECT * FROM record`)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	defer rows.Close()

	records := make([]Record, 0)
	for rows.Next() {
		record := Record{}
		err := rows.Scan(&record.Id, &record.Title, &record.Text)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return records, nil
}

func (r *RecordRepository) Read(id int64) (Record, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return Record{}, err
	}
	defer db.Close()
	row := db.QueryRow(`SELECT * FROM record WHERE id = $1`, id)
	record := Record{}
	errSc := row.Scan(&record.Id, &record.Title, &record.Text)
	if errSc != nil {
		return Record{}, errSc
	}
	return record, nil
}

func (r *RecordRepository) Update(record Record) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	_, err = db.Exec(
		`UPDATE record
		SET title = $2, text = $3
		WHERE id = $1;`,
		record.Id,
		record.Title,
		record.Text)
	return err
}

func (r *RecordRepository) Delete(id int64) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	_, err = db.Exec(
		`DELETE FROM record
		WHERE id = $1;`,
		id)
	return err
}
