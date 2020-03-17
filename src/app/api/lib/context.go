package lib

import (
	"sync"

	"pandog/infra/local/db"
	"pandog/infra/local/lib"
)

var config2 *lib.Config
var db2 *db.MySQL

var once sync.Once

func NewConfig() *lib.Config {
	once.Do(func() {
		config2 = lib.NewConfig()
	})
	return config2
}

func NewDb() *db.MySQL {
	once.Do(func() {
		dbc := config2.GetStringMapString("db")
		db2 = db.NewMySQL(dbc)
	})
	return db2
}
