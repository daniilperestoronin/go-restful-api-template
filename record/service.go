package record

type Service struct {
	rep Repository
}

func NewService(r Repository) Service {
	return Service{
		rep: r,
	}
}

func (s Service) create(record Record) (int64, error) {
	return s.rep.create(record)
}

func (s Service) readAll() ([]Record, error) {
	return s.rep.readAll()
}

func (s Service) read(id int64) (Record, error) {
	return s.rep.read(id)
}

func (s Service) update(record Record) error {
	return s.rep.update(record)
}

func (s Service) remove(id int64) error {
	return s.rep.delete(id)
}
