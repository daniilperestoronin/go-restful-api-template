package record

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Repository interface {
	create(record Record) (int64, error)
	readAll() ([]Record, error)
	read(id int64) (Record, error)
	update(record Record) error
	delete(id int64) error
}

type PgRepository struct {
	driverName     string
	dataSourceName string
}

func NewPgRepository(driverStr string, dataSourceStr string) PgRepository {
	return PgRepository{
		driverName:     driverStr,
		dataSourceName: dataSourceStr,
	}
}

func (r PgRepository) getDb() (*sql.DB, error) {
	return sql.Open(r.driverName, r.dataSourceName)
}

func (r PgRepository) create(record Record) (int64, error) {
	db, err := r.getDb()
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

func (r PgRepository) readAll() ([]Record, error) {
	db, err := r.getDb()
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

func (r PgRepository) read(id int64) (Record, error) {
	db, err := r.getDb()
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

func (r PgRepository) update(record Record) error {
	db, err := r.getDb()
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

func (r PgRepository) delete(id int64) error {
	db, err := r.getDb()
	if err != nil {
		return err
	}
	_, err = db.Exec(
		`DELETE FROM record
		WHERE id = $1;`,
		id)
	return err
}
