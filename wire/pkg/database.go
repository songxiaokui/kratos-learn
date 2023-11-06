package pkg

import "github.com/google/wire"

type Database interface {
	Query(interface{}) interface{}
}

type Mysql struct {
	Host   string
	User   string
	Passwd string
	Db     string
}

func (m *Mysql) Query(interface{}) interface{} {
	return "mysql data"
}

type Redis struct {
	Host   string
	Passwd string
	Db     int
}

func (r *Redis) Query(interface{}) interface{} {
	return "redis data"
}

func NewDB(c Config) Database {
	switch c.DB.Type {
	case "mysql":
		return &Mysql{
			Host:   c.DB.Host,
			Passwd: c.DB.Password,
			Db:     c.DB.DB,
		}
	case "redis":
		return &Redis{
			Host:   c.DB.Host,
			Passwd: c.DB.Password,
			Db:     16,
		}
	}
	return Database(nil)
}

var ProviderDbSet = wire.NewSet(NewDB)
