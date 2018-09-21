package main

type RecordService struct {
	rRep RecordRepository
}

func NewRecordService(recRep RecordRepository) RecordService {
	return RecordService{
		rRep: recRep,
	}
}

func (rs RecordService) Create(record Record) (int64, error) {
	return rs.rRep.Create(record)
}

func (rs RecordService) ReadAll() ([]Record, error) {
	return rs.rRep.ReadAll()
}

func (rs RecordService) Read(id int64) (Record, error) {
	return rs.rRep.Read(id)
}

func (rs RecordService) Update(record Record) error {
	return rs.rRep.Update(record)
}

func (rs RecordService) Remove(id int64) error {
	return rs.rRep.Delete(id)
}
