package pkg

import (
	"github.com/google/wire"
)

type Service struct {
	db  Database
	sch Schedule
}

func NewService(db Database, sched Schedule) *Service {
	return &Service{
		db:  db,
		sch: sched,
	}
}

func (s *Service) GetData() interface{} {
	return s.db.Query("")
}

func (s *Service) RunSched() {
	s.sch.Start()
}

var ProviderServiceSet = wire.NewSet(NewService)
