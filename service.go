package main

var recordRepository = NewRecordRepository()

func create(record Record) (int64, error) {
	return recordRepository.Create(record)
}

func readAll() ([]Record, error) {
	return recordRepository.ReadAll()
}

func read(id int64) (Record, error) {
	return recordRepository.Read(id)
}

func update(record Record) error {
	return recordRepository.Update(record)
}

func remove(id int64) error {
	return recordRepository.Delete(id)
}
