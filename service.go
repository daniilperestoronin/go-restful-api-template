package main

var records = map[int64]Record{
	1: {Id: 1, Title: "Title 1", Text: "Blah Blah Blah Blah 1"},
	2: {Id: 2, Title: "Title 2", Text: "Blah Blah Blah Blah 2"},
	3: {Id: 3, Title: "Title 3", Text: "Blah Blah Blah Blah 3"},
}

func getRecords() map[int64]Record {
	return records
}

func getRecord(id int64) Record {
	return records[id]
}

func createRecord(record Record) {
	records[record.Id] = record
}
